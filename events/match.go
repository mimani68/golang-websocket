package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

func matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message.Text)
	return gosf.NewSuccessMessage("Start Game")
}

func act(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message.Room)
	return gosf.NewSuccessMessage("Act")
}

func cheat(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message)
	return gosf.NewSuccessMessage("Cheat")
}

func leave(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request)
	return gosf.NewSuccessMessage("Leave")
}
