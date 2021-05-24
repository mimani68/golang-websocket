package state

import (
	"fmt"

	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/struct_helper"

	"blackoak.cloud/balout/v2/components/redis"
	"blackoak.cloud/balout/v2/model"
)

const (
	// Off StateType = "Off"
	// On  StateType = "On"

	None  StateType = "None"
	Guess StateType = "Guess"
	Cheat StateType = "Cheat"

	// SwitchOff EventType = "SwitchOff"
	// SwitchOn  EventType = "SwitchOn"

	NoneEvent  EventType = "NoneEvent"
	GuessEvent EventType = "GuessEvent"
	CheatEvent EventType = "CheatEvent"
)

type GuessWordAction struct{}

func (a *GuessWordAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("Player guess word")
	//
	// Store player new state
	//
	ctx := eventCtx.(struct {
		Id     string
		Player string
	})
	s, statePool := redis.GetKVJson("game:state:" + ctx.Id)
	if !s {
		log.Log("Redis database connection error in get 'game:state:'" + ctx.Id)
	}
	b := statePool.(struct {
		Events []interface{}
	})
	b.Events = append(b.Events, map[string]string{
		"actin":  "guess word",
		"time":   "2012",
		"player": ctx.Player,
	})
	successStore := redis.SetKV("game:state:"+ctx.Id, struct_helper.ToMap(b), 50000)
	if !successStore {
		log.Log("Failed store 'game:state:'" + ctx.Id + " in database")
	}
	return NoOp
}

type CheatAction struct{}

func (a *CheatAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("Player cheat in the game")
	return NoOp
}

// type OffAction struct{}

// func (a *OffAction) Execute(eventCtx EventContext) EventType {
// 	fmt.Println("The light has been switched off")
// 	return NoOp
// }

// type OnAction struct{}

// func (a *OnAction) Execute(eventCtx EventContext) EventType {
// 	fmt.Println("The light has been switched on")
// 	return NoOp
// }

func NewGameSwitchFSM(r model.Room) *StateMachine {
	//
	// load state from redis
	//
	// s, statePool := redis.GetKVJson("game:state:" + r.Id)
	s, _ := redis.GetKVJson("game:state:" + r.Id)
	if !s {
		log.Log("Redis database connection error in get 'game:state:'" + r.Id)
	}
	// log.Logger(statePool["id"])
	return &StateMachine{
		Id: r.Id,
		States: States{
			Default: State{
				Events: Events{
					GuessEvent: None,
				},
			},
			Guess: State{
				Action: &GuessWordAction{},
				Events: Events{
					CheatEvent: Guess,
				},
			},
			Cheat: State{
				Action: &CheatAction{},
				Events: Events{
					GuessEvent: Cheat,
				},
			},
		},
	}
}
