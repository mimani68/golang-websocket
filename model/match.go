package model

import (
	"blackoak.cloud/balout/v2/components/redis"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/config"
)

type Room struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	MaxPlayer   int    `json:"max_player,omitempty"`
	Players     []Player
	JoiningList []Player
	Groups      []Group
	Policy      []Policy
	Actions     []Action
	Envirnoment Envirnoment
	TurnHistory string `json:"trun_history,omitempty"`
	TurnCurrent string `json:"trun_current,omitempty"`
	Blocked     []Player
	Winners     []Player
	Losers      []Player
	Owner       Player
	Creator     Player
	Date        DateModel
	Status      string `json:"status,omitempty"`
}

var RoomEnum = struct {
	ACTIVE   string
	DEACTIVE string
	PENDING  string
}{
	ACTIVE:   "ACTIVE",
	DEACTIVE: "DEACTIVE",
	PENDING:  "PENDING",
}

var RoomTypeEnum = struct {
	PUBLIC  string
	PRIVATE string
}{
	PUBLIC:  "PUBLIC",
	PRIVATE: "PRIVATE",
}

func (room *Room) Store() bool {
	collectionString := "room"
	a := redis.SetKV(config.REDIS_RECORD_PREFIX+collectionString+":id:"+room.Id, room.ToMap(), config.REDIS_DATA_TTL)
	if a {
		return true
	} else {
		return false
	}
}

func (room *Room) ToMap() map[string]interface{} {
	return struct_helper.ToMap(room)
}
