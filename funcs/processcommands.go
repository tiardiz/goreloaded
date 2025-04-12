package firstproject

func ProcessCommands(words []string) []string {
	for i := 0; i < len(words); i++ {

		words = BinHexProcess(words)
		words = UpLowCapCommands(words)
		words = CommandsWithNums(words)

	}
	correctedText := CorrectArticles(words)
	return correctedText
}
