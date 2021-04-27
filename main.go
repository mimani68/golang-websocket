package main

import (
	"os"

	"github.com/ambelovsky/gosf"

	e "blackoak.cloud/balout/v2/events"
)

func init() {

	if value, exist := gosf.App.Env["GOSF_ENV"]; exist && value != "dev" {
		// Prod/Stage Config
		gosf.LoadConfig("server", "./config/server-secure.json")
	} else {
		// Default and "dev" config
		gosf.LoadConfig("server", "./config/server.json")
	}

	e.Routers()

}

func BaloutOnlineGame() {
	serverConfig := gosf.App.Config["server"].(map[string]interface{})
	if len(os.Getenv("PORT")) != 0 {
		serverConfig["PORT"] = os.Getenv("PORT")
	}
	gosf.Startup(serverConfig)
}

func main() {
	BaloutOnlineGame()
}
