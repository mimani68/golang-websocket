package events

import (
	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/helper/gosf"
)

func OnConnectHandler(client *gosf.Client, request *gosf.Request) {
	log.Log("Client connected.")
}

func OnDisconnectHandler(client *gosf.Client, request *gosf.Request) {
	log.Log("Client disconnected.")
}

func BeforeRequestHandler(client *gosf.Client, request *gosf.Request) {
	log.Log("Before request hook")
}
