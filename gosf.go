package main

import "github.com/ambelovsky/gosf"

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	println(request.Message.Text)
	return gosf.NewSuccessMessage(request.Message.Text)
}

func init() {
	// Listen on an endpoint
	gosf.Listen("ping", echo)
}

func websocket_app() {
	// Start the server using a basic configuration
	gosf.Startup(map[string]interface{}{
		"port": 3000,
		"path": "/balout/api/v1/match/",
	})
}
