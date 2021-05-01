package events

import (
	"fmt"

	isc "blackoak.cloud/balout/v2/components/interservice_communication"
	"blackoak.cloud/balout/v2/helper/gosf"
)

// type Store interface {
// 	GetByID(uint) (*model.User, error)
// 	GetByEmail(string) (*model.User, error)
// 	GetByUsername(string) (*model.User, error)
// 	Create(*model.User) error
// 	Update(*model.User) error
// 	AddFollower(user *model.User, followerID uint) error
// 	RemoveFollower(user *model.User, followerID uint) error
// 	IsFollower(userID, followerID uint) (bool, error)
// }

type Player struct{}

func (p *Player) authenticate(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// token := string(request.Message.Token)
	// get player from external server
	// if player exist
	// store token in redis and say loging: true
	// else say login: false
	a := new(isc.Http)
	result := a.Get()
	// byteSlice, err := json.Marshal(result)
	// if err != nil {
	// 	panic("marshal code.json: " + err.Error())
	// }
	return gosf.NewSuccessMessage("message goes here", result)
}

func (p *Player) playerIdentity(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// get token from redis
	// if player valid continus else reject

	fmt.Print(request.Message.Text)

	// var store = []byte("")
	// store = append(store, string("salam"))
	// store = append(store, string("mahdi"))
	// fmt.Println(store)
	// a, _ := json.Marshal(client)
	// byt := []byte(a)
	// var dat map[string]interface{}
	// type Dat struct {
	// 	Channel string `json:"channel,omitempty"`
	// 	Roome   string `json:"room,omitempty"`
	// }
	// datVar := make(Dat)
	// if err := json.Unmarshal(byt, &datVar); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(datVar)
	return gosf.NewSuccessMessage("Whoami")
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
