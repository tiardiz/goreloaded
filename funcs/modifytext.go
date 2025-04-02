package firstproject

import (
	"regexp"
	"strings"
)

func ModifytText(text string) string {
	lines := strings.Split(text, "\n")

	for i := range lines {

		addprobelposlesymbols := regexp.MustCompile(`([,.!?)}])(\S)`)
		lines[i] = addprobelposlesymbols.ReplaceAllString(lines[i], "$1 $2")
		removeprobelsperedcsymbols := regexp.MustCompile(`\s*([.,;!?)}[\]-])`)
		lines[i] = removeprobelsperedcsymbols.ReplaceAllString(lines[i], "$1")
		doposlekavychek := regexp.MustCompile(`'\s*(.*?)\s*'`)
		lines[i] = doposlekavychek.ReplaceAllString(lines[i], "'$1'")
		doposlekavychek2 := regexp.MustCompile(`"\s*(.*?)\s*"`)
		lines[i] = doposlekavychek2.ReplaceAllString(lines[i], "\"$1\"")
		probelperedskopsadd := regexp.MustCompile(`([^\s])([[{(])`)
		lines[i] = probelperedskopsadd.ReplaceAllString(lines[i], "$1 $2")
		probelperedskops := regexp.MustCompile(`\s+([{(])`)
		lines[i] = probelperedskops.ReplaceAllString(lines[i], " $1")
		probels := regexp.MustCompile(`\s+`)
		lines[i] = probels.ReplaceAllLiteralString(strings.TrimSpace(lines[i]), " ")

	}

	cleanedText := strings.Join(lines, "\n")

	return cleanedText
}
