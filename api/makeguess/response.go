package makeguess

type ResponseBody struct {
	Level    int    `json:"level"`
	Question string `json:"question"`
}
