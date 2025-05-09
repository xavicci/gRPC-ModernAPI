package main

import (
	restfibonacciserver "github.com/xavicci/gRPC-ModernAPI/grpc-fibonacci/apps/rest-fibonacci-server"
)

func main() {
	app := restfibonacciserver.NewApp()
	app.Start()
}
