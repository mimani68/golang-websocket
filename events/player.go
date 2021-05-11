package events

import (
	"blackoak.cloud/balout/v2/helper/gosf"
	model "blackoak.cloud/balout/v2/model"
)

func authenticate(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token := string(request.Message.Token)
	a := new(model.Player)
	if token == "" {
		return gosf.NewFailureMessage("Invalid Token")
	}
	profile := a.GetByToken(token)
	if profile.Id == "" {
		return gosf.NewFailureMessage("Invalid player")
	}
	profile.Store()
	return gosf.NewSuccessMessage("Welcome", profile.ToMap())
}

func playerIdentity(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage("Whoami")
}
