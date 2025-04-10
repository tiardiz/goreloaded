package firstproject

import (
	"strings"
)

func ControlProcess(words []string) []string {
	// Повторяем обработку, пока есть команды в строках
	for {
		changed := false // Флаг, который будет показывать, были ли изменения в этом проходе

		for i := 0; i < len(words); i++ {

			// Обрабатываем вложенные команды, такие как (LOW(low))
			if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
				inner := words[i][1 : len(words[i])-1] // Убираем внешние скобки
				innerWords := strings.Fields(inner)    // Разбиваем на отдельные слова (например, "low")

				// Рекурсивно обрабатываем вложенные команды
				processedInner := UpLowCapCommands(innerWords)

				// Заменяем текущий элемент на обработанный результат
				words[i] = processedInner[0]
				changed = true
			}
		}

		// Если не было изменений, значит, все команды обработаны
		if !changed {
			break
		}
	}
	return words
}
