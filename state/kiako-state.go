package state

import "fmt"

const (
	Off StateType = "Off"
	On  StateType = "On"

	SwitchOff EventType = "SwitchOff"
	SwitchOn  EventType = "SwitchOn"
)

type OffAction struct{}

func (a *OffAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("The light has been switched off")
	return NoOp
}

type OnAction struct{}

func (a *OnAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("The light has been switched on")
	return NoOp
}

func NewGameSwitchFSM() *StateMachine {
	return &StateMachine{
		States: States{
			Default: State{
				Events: Events{
					SwitchOff: Off,
				},
			},
			Off: State{
				Action: &OffAction{},
				Events: Events{
					SwitchOn: On,
				},
			},
			On: State{
				Action: &OnAction{},
				Events: Events{
					SwitchOff: Off,
				},
			},
		},
	}
}
