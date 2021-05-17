package model

type Match struct {
	Id             string `json:"id,omitempty"`
	Type           string `json:"type,omitempty"`
	MaxPlayer      int    `json:"max_player,omitempty"`
	Room           Room
	Player         []Player
	RequestForJoin []Player
	Group          []Group
	Policy         []Policy
	Action         Action
	Envirnoment    Envirnoment
	Turn_history   string
	Turn_current   string
	Blocked        []Player
	Winner         []Player
	Date           struct {
		ExpireAt string
	}
	Status string
}

type Room struct {
	Id        string   `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Child     string   `json:"child,omitempty"`
	Type      string   `json:"type,omitempty"`
	Groups    []Group  `json:"group,omitempty"`
	Players   []Player `json:"player,omitempty"`
	MaxPlayer int      `json:"max_player,omitempty"`
	Creator   Player
	Owner     Player
	Status    string `json:"status,omitempty"`
}

type Group struct {
	Id         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	MaxMemeber string `json:"max_mamber,omitempty"`
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
