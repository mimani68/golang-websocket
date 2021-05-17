package model

type RequestDto struct {
	Room string `json:"room,omitempty"`
	Word string `json:"word,omitempty"`
}
