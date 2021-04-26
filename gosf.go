package main

import "github.com/ambelovsky/gosf"

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage(request.Message.Text)
}

func init() {
	// Listen on an endpoint
	gosf.Listen("echo", echo)
}

func websocket_app() {
	// Start the server using a basic configuration
	gosf.Startup(map[string]interface{}{
		"port": 3001,
		"path": "/balout/api/v1/match/",
	})
}
