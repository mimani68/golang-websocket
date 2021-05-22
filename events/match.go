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

<<<<<<< HEAD
func matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message.Text)
	return gosf.NewSuccessMessage("Start Game")
}

func act(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message.Room)
	return gosf.NewSuccessMessage("Act")
}

func cheat(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request.Message)
	return gosf.NewSuccessMessage("Cheat")
}

func leave(client *gosf.Client, request *gosf.Request) *gosf.Message {
	fmt.Println(request)
=======
func creatNewMatch(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// 1- get player data
	me := new(model.Player)
	success, player := me.GetPlayerBySessionId(client.GetSessinId())
	if !success && player.Id == "" {
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
	room.DoAction("DO GUESS")
	// 3- join player to room
	// room.Players = append(room.Players, *me)
	room.JoinPlayerToRoom(*me)
	// 4- update redis
	s := room.Store()
	if !s {
		log.Log("[error] %w", s)
		// panic()
	}

	log.Log("[DEBUG] %s" + "The new room created")
	return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(room))
}

func matchStart(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		panic(err)
	}
	// 1- check if room is exists
	p := new(model.Player)
	success, player := p.GetPlayerBySessionId(client.GetSessinId())
	if !success && player.Id == "" {
		return gosf.NewFailureMessage("Invalid player")
	}
	// 2- join player to room
	a := new(model.Room)
	stat, r := a.GetById(req.Room)
	if !stat && r.Id == "" {
		return gosf.NewFailureMessage("Invalid room")
	}
	r.JoinPlayerToRoom(*player) // automatic join player.JoinToRoom()
	client.Join(req.Room)
	// 3- create word of game play
	words := make(map[string]interface{})
	words["word"] = "salam"
	words["count"] = "1"
	_ = r.SetWordsAssets(words)
	// 3- send event for other players
	client.Broadcast(req.Room, "balout:match:start", &gosf.Message{
		Text: "new player join room",
		Body: words,
	})
	return gosf.NewSuccessMessage("Start Game", struct_helper.ToMap(req))
}

func act(client *gosf.Client, request *gosf.Request) *gosf.Message {
	var req model.RequestDto
	err := mapstructure.Decode(request.Message.Body, &req)
	if err != nil {
		return gosf.NewFailureMessage("Invalid room")
	}
	if req.Room == "" {
		return gosf.NewFailureMessage("Invalid room")
	}
	r := new(model.Room)
	r.GetById(req.Room)
	// Check player
	p := new(model.Player)
	s, player := p.GetPlayerBySessionId(client.GetSessinId())
	if !s && player.Id == "" {
		return gosf.NewFailureMessage("Invalid player")
	}
	//
	// 1- check income word is correct form
	if req.Word == "" {
		return gosf.NewFailureMessage("Empty word")
	}
	if successAct := r.IsWordExists(req.Word); successAct {
		// 2- update data in redis
		r.DoAction("Correct Guess")
		// 3- inform other player that you play
		client.Broadcast(r.Id, "balout:match:player:act", &gosf.Message{
			Text: "Player " + player.Id + " act correctly.",
		})
		return gosf.NewFailureMessage("Empty word")
	}
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
	qq := fmt.Sprintf("%s", request.Message.Body["room"])
	if qq == "" {
		// panic(err)
		return gosf.NewFailureMessage("Empty room name")
	}
	var room model.Room
	var o model.Player
	_, player := o.GetPlayerBySessionId(client.GetSessinId())
	//
	// 1- decrease online-win of player
	player.IncreaseOnlineFailed(room.Id)
	// 2- increase opponent online-win
	for i := 0; i < len(room.Players); i++ {
		// if (room.Players[i].(interface{}))["id"] == player.Id {
		if room.Players[i].Id != player.Id {
			var o model.Player
			_, p := o.GetPlayerBySessionId(client.GetSessinId())
			p.IncreaseOnlineWin(room.Id)
		}
	}
	// 3- inform other player that you left
	var q gosf.Message
	q.Text = "Other player leave the game"
	client.Broadcast(room.Id, "", &q)
	// 4- update data in redis
	a_status := player.LeaveFromRoom(room.Id)
	if !a_status {
		return gosf.NewFailureMessage("No room register for this user")
	}
	b_status := room.RemovePlayerFromMatch(player.Id)
	if !b_status {
		return gosf.NewFailureMessage("No room register for this user")
	}
>>>>>>> 0fad54ab6d4db74b48ebd67dd3385fbf2b8ae634
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
