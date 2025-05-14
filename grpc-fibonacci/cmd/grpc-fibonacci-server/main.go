package main

import grpcfibonacciserver "github.com/xavicci/gRPC-ModernAPI/grpc-fibonacci/apps/grpc-fibonacci-server"

func main() {
	app := grpcfibonacciserver.NewApp()
	app.Start()
}
