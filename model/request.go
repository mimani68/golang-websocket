package model

type RequestDto struct {
	Room string `json:"room,omitempty"`
	Word string `json:"word,omitempty"`
}

type ActionDto struct {
	Type      string `json:"type,omitempty"`    /* type: 'fire' */
	Value     string `json:"value,omitempty"`   /* value: '1,23,12,0.1,2' */
	Payload   string `json:"payload,omitempty"` /* payload: '{}' */
	Room      string `json:"room,omitempty"`    /* room: '8345874yt' */
	IsPrivate bool   `json:"private,omitempty"` /* private: true */
}
