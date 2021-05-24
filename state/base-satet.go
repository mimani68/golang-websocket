package state

import (
	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/redis"
)

type EventContextStruct struct {
	Id     string
	Player string
}

type BaseAction struct{}

func (a *BaseAction) GetState(stateUniqueId string) interface{} {
	s, statePool := redis.GetKVJson("game:state:" + stateUniqueId)
	if !s {
		log.Log("Redis database connection error in get 'game:state:'" + stateUniqueId)
	}
	return statePool
}
