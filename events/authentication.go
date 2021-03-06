package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
	model "blackoak.cloud/balout/v2/model"
)

func authenticate(client *gosf.Client, request *gosf.Request) *gosf.Message {
	token := fmt.Sprintf("%s", request.Message.Body["token"])
	if token == "%!s(<nil>)" || token == "" || len(token) <= 10 {
		return gosf.NewFailureMessage("Empty or Invalid Token")
	}
	a := new(model.Player)
	if token == "" {
		return gosf.NewFailureMessage("Invalid Token")
	}
	_, profile := a.GetByToken(token)
	if profile.Id == "" {
		return gosf.NewFailureMessage("Invalid player")
	}
	profile.Session = client.GetSessinId()
	profile.Store()
	return gosf.NewSuccessMessage("Welcome", profile.ToMap())
}

func playerIdentity(client *gosf.Client, request *gosf.Request) *gosf.Message {
	//
	// Ooh
	//
	// a := new(gosf.Client)
	// var result = struct {
	// 	Id string
	// }{
	// 	Id: a.GetSessinId(),
	// }

	//
	// Ooh
	//
	// var result map[string]interface{}
	// result = {
	// 	"Id": "123"
	// }
	// result["Id"] = /* a.GetSessinId() */ "123"

	// result := make(map[string]interface{})
	// result["session"] = client.GetSessinId()

	player := new(model.Player)
	s, result := player.GetPlayerBySessionId(client.GetSessinId())
	if s {
		return gosf.NewSuccessMessage("Whoami", result.ToMap())
	} else {
		return gosf.NewFailureMessage("Invalid player")
	}
}
