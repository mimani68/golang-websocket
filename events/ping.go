package events

import "github.com/ambelovsky/gosf"

func Ping(client *gosf.Client, request *gosf.Request) *gosf.Message {
	println(request.Message.Text)
	// return gosf.NewSuccessMessage(request.Message.Text)
	response := new(gosf.Message)
	response.Success = true
	response.Text = request.Message.Text
	return response
}
