package events

import (
	"fmt"
	"time"

	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/config"
	"blackoak.cloud/balout/v2/helper/gosf"
	"blackoak.cloud/balout/v2/model"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func creatNewMatch(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		panic(err)
	}

	// 1- get player data
	me := new(model.Player)
	_, player := me.GetPlayerBySessionId(client.GetSessinId())
	if player.Id == "" {
		return gosf.NewFailureMessage("Invalid player")
	}
	// 2- create Room
	room := model.Room{
		Id:        uuid.New().String(),
		Type:      model.RoomTypeEnum.PRIVATE,
		MaxPlayer: config.MAX_PLAYER_IN_A_ROOM,
		Owner:     *me,
		Creator:   *me,
		Date: model.DateModel{
			StartAt:  time.Now().Format(time.RFC3339),
			ExpireAt: time.Now().Add(time.Duration(config.ROOM_EPIRE_AFTER_NS)).Format(time.RFC3339),
		},
		Status: model.RoomEnum.ACTIVE,
	}
	room.Actions = append(room.Actions, model.Action{
		Title:  "DO GAUSE",
		Status: model.ActionEnum.ACTIVE,
	})
	// 3- join player to room
	// room.Players = append(room.Players, *me)
	room.Players = append(room.Players, me.Id)
	// 4- update redis
	s := room.Store()
	if !s {
		log.Log("[error] %w", s)
		// panic()
	}
	w := new(model.Player)
	w.AssignPlayerToRoom(room.Id)
	log.Log("[DEBUG] %s" + "The new room created")
	return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(room))
}

func matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		panic(err)
	}
	// ======== debug =====================================
	log.Log("[MATCH-START] room: " + req.Room)
	// fmt.Printf("room: %s", result.Room) // failed to load
	// return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(result))
	// ======== debug =====================================
	//
	// 1- check if room is exists
	// 2- join player to room
	// 3- update redis
	// 4- send event for other players
	//
	return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(req))
}

func act(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		panic(err)
	}
	log.Log("[ACT] room: " + req.Room)
	log.Log("[ACT] word: " + req.Word)
	//
	// 1- check income word is correct form
	// 2- inform other player that you play
	// 3- update data in redis
	//
	return gosf.NewSuccessMessage("Act", struct_helper.ToMap(req))
}

func cheat(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		panic(err)
	}
	log.Log("[CHEAT] room: " + req.Room)
	//
	// 1- reduse charge
	// 2- inform other player that you play
	// 3- update data in redis
	//
	return gosf.NewSuccessMessage("Cheat", struct_helper.ToMap(req))
}

func leaveSingleRoom(client *gosf.Client, request *gosf.Request) *gosf.Message {
	room := fmt.Sprintf("%s", request.Message.Body["room"])
	if room == "" {
		// panic(err)
		return gosf.NewFailureMessage("Empty room name")
	}
	a := new(model.Player)
	b := a.GetRoomOfPlayer(room)
	if len(b) < 1 {
		return gosf.NewFailureMessage("No room register for this user")
	}
	//
	// 1- decrease online-win of player
	// 2- increase opponent online-win
	// 3- inform other player that you left
	// 4- update data in redis
	//
	// return gosf.NewSuccessMessage("Leave", struct_helper.ToMap(req))
	log.Log("[LEAVE] room: " + room)
	return gosf.NewSuccessMessage("Leave")
}

func leaveAllRoom(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		panic(err)
	}
	log.Log("[LEAVE-ALL] room: " + req.Room)
	//
	// 1- decrease online-win of player
	// 2- increase opponent online-win
	// 3- inform other player that you left
	// 4- update data in redis
	//
	return gosf.NewSuccessMessage("Leave all rooms", struct_helper.ToMap(req))
}
