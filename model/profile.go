package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"

	str "blackoak.cloud/balout/v2/components/struct_helper"
	// isc "blackoak.cloud/balout/v2/components/interservice_communication"
	redis "blackoak.cloud/balout/v2/components/redis"
	cnf "blackoak.cloud/balout/v2/config"
)

type Player struct {
	Id           string `json:"id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	LoginPlayer  string `json:"loginPlayer,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
	Session      string `json:"session,omitempty"`
}

var collectionString = "player"

func (p *Player) GetByToken(token string) *Player {
	hasCache, data := redis.GetKVJson(cnf.REDIS_RECORD_PREFIX + collectionString + ":token:" + p.shorternTokenGenerator(token))
	if hasCache {
		return data.(*Player)
	} else {
		//
		// FIXME: change to real mode
		//
		// res, _ := http.Get("http://auth:3000/player/skbq623ihr359")
		// data, _ := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// return data.(*Player)
		//
		// Fake REsponse
		//
		data := &Player{
			Id:           uuid.New().String(),
			Nickname:     gofakeit.Username(),
			LoginPlayer:  uuid.New().String(),
			AccessToken:  gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
			RefreshToken: gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
			DeviceId:     gofakeit.MacAddress(),
		}
		return data
	}
}

func (p *Player) GetById(playerId string) *Player {
	hasCache, data := redis.GetKVJson(cnf.REDIS_RECORD_PREFIX + collectionString + ":" + playerId)
	if hasCache {
		return data.(*Player)
	} else {
		//
		// FIXME: change to real mode
		//
		// res, _ := http.Get("http://auth:3000/player/skbq623ihr359")
		// data, _ := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// return data.(*Player)
		//
		// Fake REsponse
		//
		data := &Player{
			Id:           uuid.New().String(),
			Nickname:     gofakeit.Username(),
			LoginPlayer:  uuid.New().String(),
			AccessToken:  gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
			RefreshToken: gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
			DeviceId:     gofakeit.MacAddress(),
		}
		return data
	}
}

func (p *Player) GetPlayerBySessionId(playerId string) (bool, *Player) {
	s, a := redis.GetKVJson(cnf.REDIS_RECORD_PREFIX + collectionString + ":" + p.Id)
	if s && fmt.Sprintf("%T", a) == "string" {
		var b interface{}
		json.Unmarshal([]byte(a.(string)), b)
		return s, b.(*Player)
	}
	if !s {
		//
		// Yeeh
		// https://stackoverflow.com/questions/50697914/return-nil-for-a-struct-in-go
		//
		return s, &Player{}
	}
	return s, a.(*Player)
}

func (p *Player) Store() bool {
	//
	// store with token id
	//
	storeOneStatus := redis.SetKV(cnf.REDIS_RECORD_PREFIX+collectionString+":token:"+p.shorternTokenGenerator(), p, 30)
	//
	//store with player id
	//
	storeTwoStatus := redis.SetKV(cnf.REDIS_RECORD_PREFIX+collectionString+":"+p.Id, p, 30)
	if storeOneStatus && storeTwoStatus {
		return true
	} else {
		return false
	}
}

// optional argument is invalid here
// A nice way to achieve something like optional parameters is to use variadic args.
// The function actually receives a slice of whatever type you specify.
func (p *Player) shorternTokenGenerator(params ...string) string {
	if len(params) <= 0 {
		result := strings.Split(p.AccessToken, " ")
		return result[1][:6]
	}
	result := strings.Split(params[0], " ")
	return result[1][:6]
}

func (p *Player) ToMap() interface{} {
	return str.ToMap(p)
}

// func (p *Player) CustomJSON(code int, i interface{}, f string) (err error) {
// 	if c.Context.Echo().Debug {
// 		return c.JSONPretty(code, i, "  ")
// 	}
// 	b, err := json.MarshalFilter(i, f)
// 	if err != nil {
// 		return
// 	}
// 	return c.JSONBlob(code, b)
// }
