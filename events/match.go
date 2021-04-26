package events

import (
	"fmt"

	"github.com/ambelovsky/gosf"
)

func matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message.Text)
	return gosf.NewSuccessMessage("Hello")
}
