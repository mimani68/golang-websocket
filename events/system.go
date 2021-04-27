package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

func OnConnectHandler(client *gosf.Client, request *gosf.Request) {
	fmt.Println("Client connected.")
}

func OnDisconnectHandler(client *gosf.Client, request *gosf.Request) {
	fmt.Println("Client disconnected.")
}

func BeforeRequestHandler(client *gosf.Client, request *gosf.Request) {
	fmt.Println("Before request hook")
}
