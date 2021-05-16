package events

import (
	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/helper/gosf"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type RequestDto struct {
	Room string `json:"room,omitempty"`
	Word string `json:"word,omitempty"`
}

func creatNewMatch(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var result RequestDto
	err := mapstructure.Decode(request.Message.Body, &result)
	if err != nil {
		panic(err)
	}
	result.Room = uuid.New().String()
	// ======== debug =====================================
	log.Log("[CREATE-NEW-MATCH] room: " + result.Room)
	// fmt.Printf("room: %s", result.Room) // failed to load
	// return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(result))
	// ======== debug =====================================
	//
	// 1- check if room is exists
	// 2- join player to room
	// 3- update redis
	// 4- send event for other players
	//
	return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(result))
}

func matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var result RequestDto
	err := mapstructure.Decode(request.Message.Body, &result)
	if err != nil {
		panic(err)
	}
	// ======== debug =====================================
	log.Log("[MATCH-START] room: " + result.Room)
	// fmt.Printf("room: %s", result.Room) // failed to load
	// return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(result))
	// ======== debug =====================================
	//
	// 1- check if room is exists
	// 2- join player to room
	// 3- update redis
	// 4- send event for other players
	//
	return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(result))
}

func act(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var result RequestDto
	err := mapstructure.Decode(request.Message.Body, &result)
	if err != nil {
		panic(err)
	}
	log.Log("[ACT] room: " + result.Room)
	log.Log("[ACT] word: " + result.Word)
	//
	// 1- check income word is correct form
	// 2- inform other player that you play
	// 3- update data in redis
	//
	return gosf.NewSuccessMessage("Act", struct_helper.ToMap(result))
}

func cheat(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var result RequestDto
	err := mapstructure.Decode(request.Message.Body, &result)
	if err != nil {
		panic(err)
	}
	log.Log("[CHEAT] room: " + result.Room)
	//
	// 1- reduse charge
	// 2- inform other player that you play
	// 3- update data in redis
	//
	return gosf.NewSuccessMessage("Cheat", struct_helper.ToMap(result))
}

func leave(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var result RequestDto
	err := mapstructure.Decode(request.Message.Body, &result)
	if err != nil {
		panic(err)
	}
	log.Log("[CHEAT] room: " + result.Room)
	//
	// 1- decrease online-win of player
	// 2- increase opponent online-win
	// 3- inform other player that you left
	// 4- update data in redis
	//
	return gosf.NewSuccessMessage("Leave", struct_helper.ToMap(result))
}
