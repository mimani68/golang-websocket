package events

import (
	"github.com/ambelovsky/gosf"
)

func ping(client *gosf.Client, request *gosf.Request) *gosf.Message {
	//
	//
	//
	response := new(gosf.Message)
	response.Text = "Pong"
	//
	// Broadcast
	//
	client.Broadcast("balout:system:ping", "example", response)
	return response
	//
	// Online message
	//
	// return gosf.NewSuccessMessage(request.Message.Text + " alo")
	//
	//
	// response := new(gosf.Message)
	// response.Success = true
	// response.Text = request.Message.Text
	// return response

}
