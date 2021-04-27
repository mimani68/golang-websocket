package events

import (
	"encoding/json"
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

// var store = make([]string, 0)

type Player struct{}

func (controller Player) playerIdentity(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Print(request.Message.Text)

	// var store = []byte("")
	// store = append(store, string("salam"))
	// store = append(store, string("mahdi"))
	// fmt.Println(store)
	a, _ := json.Marshal(client)
	byt := []byte(a)
	// var dat map[string]interface{}
	type Dat struct {
		Channel string `json:"channel,omitempty"`
		Roome   string `json:"room,omitempty"`
	}
	datVar := make(Dat)
	if err := json.Unmarshal(byt, &datVar); err != nil {
		panic(err)
	}
	fmt.Println(datVar)
	return gosf.NewSuccessMessage("Whoami")
}

func (controller Player) authenticatePlayer(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage("Whoami")
}
