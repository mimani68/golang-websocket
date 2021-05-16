package events

import (
	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/helper/gosf"
	"github.com/mitchellh/mapstructure"
)

type RequestDto struct {
	Room string `json:"room,omitempty"`
	Word string `json:"word,omitempty"`
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
	// 3- send event for other players
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
	// 3-
	//
	return gosf.NewSuccessMessage("Act")
}

func cheat(client *gosf.Client, request *gosf.Request) *gosf.Message {
	log.Log(request.Message)
	return gosf.NewSuccessMessage("Cheat")
}

func leave(client *gosf.Client, request *gosf.Request) *gosf.Message {
	log.Log(request)
	return gosf.NewSuccessMessage("Leave")
}
