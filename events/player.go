package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

var store = make([]string, 0)

type Player struct{}

func (controller Player) playerIdentity(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message.Text)

	// var store = []byte("")
	store = append(store, string("salam"))
	store = append(store, string("mahdi"))
	fmt.Println(store)
	return gosf.NewSuccessMessage("Whoami")
}

// func (controller Player) playerDetails(client *gosf.Client, request *gosf.Request) *gosf.Message {
// 	fmt.Print(request.Message.Text)
// 	return gosf.NewSuccessMessage("Hello")
// }
