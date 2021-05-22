package model

type Group struct {
	Id         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	MaxMemeber string `json:"max_member,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Member     string `json:"member,omitempty"`
	Status     string `json:"status,omitempty"`
}

type Policy struct {
	Id      string      `json:"id,omitempty"`
	Title   string      `json:"title,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type Envirnoment struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

type Action struct {
	Id      string      `json:"id,omitempty"`
	Title   string      `json:"title,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Status  string      `json:"status,omitempty"`
}

var ActionEnum = struct {
	ACTIVE   string
	DEACTIVE string
	PENDING  string
}{
	ACTIVE:   "ACTIVE",
	DEACTIVE: "DEACTIVE",
	PENDING:  "PENDING",
}

type DateModel struct {
	StartAt  string `json:"start_at,omitempty"`
	ExpireAt string `json:"expire_at,omitempty"`
}
