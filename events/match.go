package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

type Match struct{}

func (controller Match) matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message.Text)
	return gosf.NewSuccessMessage("Start Game")
}

func (controller Match) act(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message.Body["room"])
	return gosf.NewSuccessMessage("Act")
}

func (controller Match) cheat(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message)
	return gosf.NewSuccessMessage("Cheat")
}

func (controller Match) leave(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request)
	return gosf.NewSuccessMessage("Leave")
}
