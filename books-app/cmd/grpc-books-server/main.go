package main

import (
	grpcbooksserver "github.com/xavicci/gRPC-ModernAPI/books-app/internal/apps/grpc-books-server"
)

func main() {
	app := grpcbooksserver.NewBookServer()
	app.Start()

	app.Shutdown()
}
