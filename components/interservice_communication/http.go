package interservice_communication

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type Http struct{}

func (controller Http) Get() interface{} {
	a := struct {
		Id           string `json:"id,omitempty"`
		Nickname     string `json:"nickname,omitempty"`
		LoginProfile string `json:"loginProfile,omitempty"`
		AccessToken  string `json:"accessToken,omitempty"`
		RefreshToken string `json:"refreshToken,omitempty"`
		DeviceId     string `json:"deviceId,omitempty"`
	}{
		Id:           uuid.New().String(),
		Nickname:     gofakeit.Username(),
		LoginProfile: uuid.New().String(),
		AccessToken:  gofakeit.Regex(`^[az-AZ]{0,16}$`),
		RefreshToken: gofakeit.Regex(`^[az-AZ]{0,16}$`),
		DeviceId:     gofakeit.MacAddress(),
	}
	return a
}
