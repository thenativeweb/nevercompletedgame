package startgame

type ResponseBody struct {
	GameID   string `json:"gameId"`
	Level    int    `json:"level"`
	Question string `json:"question"`
}
