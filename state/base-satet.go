package state

import (
	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/redis"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/config"
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

func (a *BaseAction) SetState(stateUniqueId string, playerId string, actionType string) bool {
	statePool := a.GetState(stateUniqueId)
	b := statePool.(struct {
		Id     string
		Events []interface{}
	})
	b.Events = append(b.Events, map[string]string{
		"actin":  actionType,
		"time":   "2012",
		"player": playerId,
	})
	successStore := redis.SetKV("game:state:"+stateUniqueId, struct_helper.ToMap(b), config.REDIS_DATA_TTL)
	if !successStore {
		log.Log("Failed store 'game:state:'" + stateUniqueId + " in database")
	}
	return true
}
