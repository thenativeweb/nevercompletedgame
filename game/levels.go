package game

type level struct {
	Question string
	Answer   string
}

var levels = []level{
	{
		Question: "First letter of the alphabet.",
		Answer:   "A",
	}, {
		Question: "2 + 1 = ?",
		Answer:   "3",
	}, {
		Question: "A, B, C, D, E, _?",
		Answer:   "F",
	},
}
