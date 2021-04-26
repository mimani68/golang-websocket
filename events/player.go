package events

import (
	"fmt"

	"github.com/ambelovsky/gosf"
)

func playerIdentity(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message.Text)
	return gosf.NewSuccessMessage("Hello")
}
