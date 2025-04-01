package firstproject

import (
	"regexp"
	"strings"
)

func ModifytText(text string) string {
	lines := strings.Split(text, "\n")

	// Обрабатываем каждую строку
	for i, line := range lines {
		doposlekavychek := regexp.MustCompile(`(['"])\s(.*?)\s(['"])`)
		lines[i] = doposlekavychek.ReplaceAllString(lines[i], "$2$3")
		probelsperedcsymbols := regexp.MustCompile(`\s([.,;!?(){}[\]-])`)
		lines[i] = probelsperedcsymbols.ReplaceAllString(line, "$1")
		probels := regexp.MustCompile(`\s+`)
		lines[i] = probels.ReplaceAllLiteralString(strings.TrimSpace(lines[i]), " ")

	}

	// Объединяем строки обратно в текст с сохранением символов новой строки
	cleanedText := strings.Join(lines, "\n")

	return cleanedText
}
