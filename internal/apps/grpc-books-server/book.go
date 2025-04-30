package grpcbooksserver

import (
	"fmt"
	"log"
	"net"

	"moderngrpc/internal/pkg/configs"
	"moderngrpc/internal/pkg/db"
	"moderngrpc/internal/pkg/db/migrations"
	"moderngrpc/internal/pkg/proto"
	"moderngrpc/internal/pkg/repo"
	"moderngrpc/internal/pkg/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Server struct {
	dbConn *gorm.DB
}

func NewBookServer() *Server {
	return &Server{}
}

func (a *Server) Start() {

	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.ProvideDBConn(&appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	a.dbConn = dbConn

	repo := repo.GetNewBookRepository(a.dbConn)
	server := service.NewApp(repo)

	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, dbConn)
	if err != nil {
		log.Fatal(err)
	}

	migrator.RunMigrations()

	servAddr := fmt.Sprintf("0.0.0.0:%d", appConfig.ServerConfig.Port)

	fmt.Println("starting books gRPC server at", servAddr)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	proto.RegisterBookServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *Server) Shutdown() {
	dbInstance, _ := a.dbConn.DB()
	_ = dbInstance.Close()
}
