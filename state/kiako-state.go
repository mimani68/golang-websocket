package state

import (
	"fmt"
	"time"
)

// Singleton (use global variable)
var StatePool []StateMachine

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

//
//  FIXME: Implement state machine
// ---------------------------------------------------------
//  [*] get state from redis if state was exists
//	[*] update new state
//	[*] serialize state
//	[*] store stat of room in redis
//
func GameStateMachine(uniqueStateId string) (bool, *StateMachine) {
	var s *StateMachine
	if uniqueStateId == "" {
		return false, &StateMachine{}
	}
	ImportState(uniqueStateId)
	// Singleton
	if s.Id == "" {
		s = &StateMachine{
			Id:           uniqueStateId,
			LastInitDate: time.Now().Format(time.RFC3339),
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
	return true, s
}

func ImportState(uniqueStateId string) (bool, StateMachine) {
	var s bool
	var result StateMachine
	for i := 0; i < len(StatePool); i++ {
		//
		// Remove old state
		//
		hour, minutes, seconds := 1, 0, 0
		nowPlusOneHour := time.Now().Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds))
		if StatePool[i].LastInitDate >= nowPlusOneHour.Format(time.RFC3339) {
			fmt.Println("Remove data")
		}
		//
		// Get data from redis
		//
		if StatePool[i].Id != uniqueStateId {
			stateRetriveFromRedis := StateMachine{}
			if stateRetriveFromRedis.Id == "" {
				s = false
				result = StateMachine{}
			} else {
				//
				// deSerializeState
				//
				s = true
				// result = deSerialize(stateRetriveFromRedis)
				result = stateRetriveFromRedis
			}
		} else {
			s = false
			result = StatePool[i]
		}
	}
	return s, result
}

func ExportState() bool {
	return true
}
