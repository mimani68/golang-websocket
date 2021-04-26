package events

import (
	"fmt"

	"github.com/ambelovsky/gosf"
)

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message)
	return gosf.NewSuccessMessage("ECHO ")
}
