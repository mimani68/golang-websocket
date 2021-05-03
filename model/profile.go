package model

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"

	str "blackoak.cloud/balout/v2/components/struct_helper"
	// isc "blackoak.cloud/balout/v2/components/interservice_communication"
	redis "blackoak.cloud/balout/v2/components/redis"
)

type Profile struct {
	Id           string `json:"id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	LoginProfile string `json:"loginProfile,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
}

func (p *Profile) GetByToken(token string) *Profile {
	p = &Profile{
		Id:           uuid.New().String(),
		Nickname:     gofakeit.Username(),
		LoginProfile: uuid.New().String(),
		AccessToken:  gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
		RefreshToken: gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
		DeviceId:     gofakeit.MacAddress(),
	}
	return p
}

func (p *Profile) ToMap() interface{} {
	return str.ToMap(p)
}

func (p *Profile) Store() bool {
	a := redis.RedisClient()
	fmt.Print(a)
	b := redis.SetKV("profile:"+p.Id, p, 30)
	fmt.Print(b)
	return false
}
