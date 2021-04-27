package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

func dev(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message.Text)
	return gosf.NewSuccessMessage("Hello")
}
