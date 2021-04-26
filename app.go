package main

import (
	e "blackoak.cloud/balout/v2/events"
	gosf "github.com/ambelovsky/gosf"
)

func init() {

	if value, exist := gosf.App.Env["GOSF_ENV"]; exist && value != "dev" {
		// Prod/Stage Config
		gosf.LoadConfig("server", "./config/server-secure.json")
	} else {
		// Default and "dev" config
		gosf.LoadConfig("server", "./config/server.json")
	}

	e.EventList()
}

func websocket_app() {
	serverConfig := gosf.App.Config["server"].(map[string]interface{})
	gosf.Startup(serverConfig)
}
