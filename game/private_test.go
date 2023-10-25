package game

func Questions() []string {
	questions := make([]string, len(levels))

	for _, level := range levels {
		questions = append(questions, level.Question)
	}

	return questions
}

func QuestionForLevel(level int) string {
	return levels[level-1].Question
}

func Answers() []string {
	answers := make([]string, len(levels))

	for _, level := range levels {
		answers = append(answers, level.Answer)
	}

	return answers
}

func AnswerForLevel(level int) string {
	return levels[level-1].Answer
}
