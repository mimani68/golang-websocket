package state

import (
	"fmt"

	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/components/struct_helper"
	"blackoak.cloud/balout/v2/config"

	"blackoak.cloud/balout/v2/components/redis"
	"blackoak.cloud/balout/v2/model"
)

func (a *BaseAction) SetState(stateUniqueId string, playerId string, actionType string) bool {
	statePool := a.GetState(stateUniqueId)
	b := statePool.(struct {
		Id     string
		Events []interface{}
	})
	b.Events = append(b.Events, map[string]string{
		"actin":  actionType,
		"time":   "2012",
		"player": playerId,
	})
	successStore := redis.SetKV("game:state:"+stateUniqueId, struct_helper.ToMap(b), config.REDIS_DATA_TTL)
	if !successStore {
		log.Log("Failed store 'game:state:'" + stateUniqueId + " in database")
	}
	return true
}

const (
	None  StateType = "None"
	Guess StateType = "Guess"
	Cheat StateType = "Cheat"

	NoneEvent  EventType = "NoneEvent"
	GuessEvent EventType = "GuessEvent"
	CheatEvent EventType = "CheatEvent"
)

type GuessWordAction struct {
	BaseAction
}

func (a *GuessWordAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("Player guess word")
	ctx := eventCtx.(EventContextStruct)
	a.SetState(ctx.Id, ctx.Player, "Guess word")
	return NoOp
}

type CheatAction struct {
	BaseAction
}

func (a *CheatAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("Player cheat in the game")
	ctx := eventCtx.(EventContextStruct)
	a.SetState(ctx.Id, ctx.Player, "Cheat word")
	return NoOp
}

func NewGameSwitchFSM(r model.Room) *StateMachine {
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
