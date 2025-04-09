package firstproject

import (
	"strings"
)

func ProcessText(text string) string {

	lines := strings.Split(text, "\n")

	for i, line := range lines {
		line = Punctuation(line) // обработка пунктуации
		words := strings.Fields(line)
		words = MergeStrings(words)     // merge "(cap," + "<number>)", также отделяет от остальных
		words = ProcessCommands(words)  //(up), (cap), (low), (bin), (hex) functions then articles
		words = CommandsWithNums(words) // (cap|up|low, <numbers>)
		//words = HexBinProcess(words)    //binhex commands

		lines[i] = strings.Join(words, " ")
		lines[i] = Punctuation(lines[i]) //контрольная коррекция пунктуации после команд
	}

	result := strings.Join(lines, "\n")

	return result
}
