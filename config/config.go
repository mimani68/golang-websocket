package config

import "os"

var (
	REDIS_URL           = os.Getenv("REDIS_URL")
	REDIS_PORT          = os.Getenv("REDIS_PORT")
	REDIS_PASSWORD      = os.Getenv("REDIS_PASSWORD")
	REDIS_RECORD_PREFIX = "balout:"
	REDIS_DATA_TTL      = 30 * 100

	MAX_PLAYER_IN_A_ROOM = 2
	ROOM_EPIRE_AFTER_NS  = 10 * 100000
)
