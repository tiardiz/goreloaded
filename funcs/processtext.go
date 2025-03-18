package firstproject

import (
	"strings"
)

func ProcessText(text string) string {
	// Разделяем текст на строки
	lines := strings.Split(text, "\n")

	// Обрабатываем каждую строку
	for i, line := range lines {
		// Разделяем строку на слова по пробелам
		words := strings.Fields(line)

		// Обрабатываем слова (например, применяем ProcessCommands)
		words = ProcessCommands(words)

		// Собираем строку обратно, разделяя слова пробелами
		lines[i] = strings.Join(words, " ")
	}

	// Собираем текст обратно, сохраняя новые строки
	result := strings.Join(lines, "\n")

	return result
}
