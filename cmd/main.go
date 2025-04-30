package main

import (
	grpcbooksserver "moderngrpc/internal/apps/grpc-books-server"
)

func main() {
	app := grpcbooksserver.NewBookServer()
	app.Start()

	app.Shutdown()
}
