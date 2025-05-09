package main

import restfibonacciclient "github.com/xavicci/gRPC-ModernAPI/grpc-fibonacci/apps/rest-fibonacci-client"

func main() {
	app := restfibonacciclient.NewApp()
	app.Start()
}
