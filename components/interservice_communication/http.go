package interservice_communication

import (
	str "blackoak.cloud/balout/v2/components/struct_helper"
	model "blackoak.cloud/balout/v2/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type Http struct {
	dev  bool
	Data model.Profile
}

func (h *Http) Get() *Http {
	h.Data = model.Profile{
		Id:           uuid.New().String(),
		Nickname:     gofakeit.Username(),
		LoginProfile: uuid.New().String(),
		AccessToken:  gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
		RefreshToken: gofakeit.Regex(`jwt [a-zA-Z]{128}$`),
		DeviceId:     gofakeit.MacAddress(),
	}
	h.dev = false
	return h
}

func (h *Http) ToMap() interface{} {
	return str.ToMap(h.Data)
}
