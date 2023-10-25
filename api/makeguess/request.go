package makeguess

type RequestBody struct {
	GameID string `json:"gameId"`
	Guess  string `json:"guess"`
}
