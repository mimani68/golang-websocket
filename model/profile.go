package model

import (
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

func (p *Player) GetPlayerBySessionId(sessionId string) (bool, *Player) {
	s, a := redis.GetKVJson(cnf.REDIS_RECORD_PREFIX + collectionString + ":session:" + sessionId)
	var b interface{}
	// mapstructure.Decode(b, &a)
	b = str.ToStruct(a)
	// if s && fmt.Sprintf("%T", a) == "string" {
	// 	var b interface{}
	// 	_ = json.Unmarshal([]byte(a.(string)), b)
	// 	return s, b.(*Player)
	// }
	if !s {
		//
		// Yeeh
		// https://stackoverflow.com/questions/50697914/return-nil-for-a-struct-in-go
		//
		return s, &Player{}
	}
	return s, b.(*Player)
}

func (p *Player) Store() bool {
	//
	// store with token id
	//
	storeOneStatus := redis.SetKV(cnf.REDIS_RECORD_PREFIX+collectionString+":token:"+p.shorternTokenGenerator(), p.ToMap(), cnf.REDIS_DATA_TTL)
	//
	// store with player id
	//
	storeTwoStatus := redis.SetKV(cnf.REDIS_RECORD_PREFIX+collectionString+":id:"+p.Id, p.ToMap(), cnf.REDIS_DATA_TTL)
	//
	// store with Session id
	//
	storeThreeStatus := redis.SetKV(cnf.REDIS_RECORD_PREFIX+collectionString+":session:"+p.Session, p.ToMap(), cnf.REDIS_DATA_TTL)
	if storeOneStatus && storeTwoStatus && storeThreeStatus {
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
		return result[1][:10]
	}
	result := strings.Split(params[0], " ")
	return result[1][:10]
}

func (p *Player) ToMap() map[string]interface{} {
	return str.ToMap(p)
}

func (p *Player) ToStruct() struct{} {
	return str.ToStruct(p)
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
