# PASOS

docker-compose up
go run .\books-app\cmd\grpc-books-server\main.go -configFile .\books-app\configs\grpc-books-server.yaml
protoc --go_out=books-app/internal/pkg/proto --go-grpc_out=books-app books-app/internal/pkg/proto/book.proto
