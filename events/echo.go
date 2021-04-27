package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message)
	return gosf.NewSuccessMessage("ECHO ")
}
