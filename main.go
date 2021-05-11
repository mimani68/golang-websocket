package main

import (
	"os"
	"strconv"

	// "github.com/ambelovsky/gosf"
	"blackoak.cloud/balout/v2/components/redis"
	"blackoak.cloud/balout/v2/helper/gosf"

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
	var port, _ = strconv.Atoi(os.Getenv("PORT"))
	if len(os.Getenv("PORT")) != 0 {
		serverConfig["port"] = port
	}
	gosf.Startup(serverConfig)
}

func main() {
	redis.RedisClient()
	BaloutOnlineGame()
}
