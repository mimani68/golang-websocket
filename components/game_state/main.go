package gamestate

type configFile map[string]interface{}

func GetGameState(a string) (b configFile) {
	b = configFile{
		"id":                 "345",
		"game":               a,
		"max_player_in_room": "2",
	}
	return
}
