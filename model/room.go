package model

import (
	"blackoak.cloud/balout/v2/components/redis"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/config"
	"github.com/mitchellh/mapstructure"
)

type Room struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	MaxPlayer   int    `json:"max_player,omitempty"`
	Game        string
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
	Words       map[string]interface{}
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

func (r *Room) GetById(roomId string) (bool, *Room) {
	collectionString := "room"
	s, data := redis.GetKVJson(config.REDIS_RECORD_PREFIX + collectionString + ":" + roomId)
	r.importFromInterface(data)
	if s {
		return s, r
	} else {
		//
		// FIXME: change to real mode
		//
		// res, _ := http.Get("http://auth:3000/player/skbq623ihr359")
		// data, _ := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// return data.(*Player)
		// if res == nil {
		// 	return false, &Player{}
		// }
		//
		// Fake Response
		//
		// a := &Room{
		// 	Id:           uuid.New().String(),
		// 	Nickname:     gofakeit.Username(),
		// 	LoginPlayer:  uuid.New().String(),
		// 	AccessToken:  gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
		// 	RefreshToken: gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
		// 	DeviceId:     gofakeit.MacAddress(),
		// }
		return false, &Room{}
	}
}

// FIXME:
func (r *Room) RemovePlayerFromMatch(userId string) bool {
	// collectionString := "room"
	// a := redis.SetKV(config.REDIS_RECORD_PREFIX+collectionString+":id:"+room.Id, room.ToMap(), config.REDIS_DATA_TTL)
	// // a := redis.RunCommand("FT.SEARCH ")
	// if a {
	// 	return Room{}
	// } else {
	// 	return Room{}
	// }
	return true
}

func (r *Room) DoAction(actionTitle string) bool {
	r.Actions = append(r.Actions, Action{
		Title:  actionTitle,
		Status: ActionEnum.ACTIVE,
	})
	return true
}

func (r *Room) JoinPlayerToRoom(p Player) bool {
	r.Players = append(r.Players, p)
	w := new(Player)
	_, a := w.GetById(p.Id)
	a.AssignPlayerToRoom(r.Id)
	return true
}

func (r *Room) SetWordsAssets(words map[string]interface{}) bool {
	r.Words = words
	return r.Store()
}

func (r *Room) IsWordExists(word string) bool {
	for _, v := range r.Words {
		if v == word {
			return true
		}
	}
	return false
}

func (r *Room) ToMap() map[string]interface{} {
	return struct_helper.ToMap(r)
}

func (r *Room) importFromInterface(input interface{}) {
	mapstructure.Decode(input, &r)
}
