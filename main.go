package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ProcessText обрабатывает текст, сохраняя новые строки
func ProcessText(text string) string {
	// Шаг 1: Запомнить позиции новых строк в исходном тексте
	var newLineIndices []int
	for i, char := range text {
		if char == '\n' {
			newLineIndices = append(newLineIndices, i)
		}
	}

	// Шаг 2: Убрать новые строки
	textWithoutNewlines := strings.ReplaceAll(text, "\n", " ")

	// Шаг 3: Обработать текст
	words := strings.Fields(textWithoutNewlines)
	words = ProcessCommands(words)

	// Шаг 4: Собрать текст обратно
	result := strings.Join(words, " ")

	// Шаг 5: Восстановить новые строки
	resultRunes := []rune(result)
	for _, index := range newLineIndices {
		// Если индекс находится в пределах длины текста
		if index < len(resultRunes) {
			// Вставляем символ новой строки
			resultRunes = append(resultRunes[:index], append([]rune{'\n'}, resultRunes[index:]...)...)
		}
	}

	return string(resultRunes)
}

// ProcessCommands обрабатывает команды в срезе слов
func ProcessCommands(words []string) []string {
	for i := 0; i < len(words); i++ {
		// Обработка (up, <число>)
		if strings.HasPrefix(words[i], "(up,") {
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				numStr := strings.TrimSuffix(words[i+1], ")")
				num, err := strconv.Atoi(numStr)
				if err == nil && i-num >= 0 {
					for j := i - num; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}
		}

		// Обработка (cap)
		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = Capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		// Обработка (cap, <число>)
		if strings.HasPrefix(words[i], "(cap,") {
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				numStr := strings.TrimSuffix(words[i+1], ")")
				num, err := strconv.Atoi(numStr)
				if err == nil && i-num >= 0 {
					for j := i - num; j < i; j++ {
						words[j] = Capitalize(words[j])
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}
		}
	}
	return words
}

// Capitalize переписывает слово с заглавной буквы
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	// Переводим первый символ в верхний регистр
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func main() {
	// Проверка аргументов командной строки
	if len(os.Args) != 3 {
		fmt.Println("Использование: go run main.go <входной_файл> <выходной_файл>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Чтение файла
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	text := string(content)

	// Обработка текста
	processedText := ProcessText(text)

	// Запись результата в файл
	err = os.WriteFile(outputFile, []byte(processedText), 0o644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Обработка завершена. Результат сохранен в", outputFile)
}

// Разбиваем текст на слова по пробелам
// words := strings.Fields(text) // Fields автоматически удаляет лишние пробелы
// // Проходим по словам и ищем (cap)
// for i := 0; i < len(words)-1; i++ {
// 	if words[i+1] == "(cap)" {
// 		words[i] = Capitalize(words[i])
// 	}
// 	if words[i] == "(cap)" {
// 		words[i] = strings.TrimPrefix(words[i], "(cap)")
// 	}
// 	if strings.HasPrefix(words[i], "(cap,") {
// 		// Извлекаем число из следующего элемента
// 		if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
// 			// Убираем ")" из числа
// 			numStr := strings.TrimSuffix(words[i+1], ")")
// 			num, err := strconv.Atoi(numStr)
// 			if err != nil {
// 				// Если число не распознано, пропускаем
// 				continue
// 			}

// 			// Проверяем, что количество слов для обработки не превышает доступных
// 			if i-num >= 0 {
// 				// Переписываем указанное количество слов с заглавной буквы
// 				for j := i - num; j < i; j++ {
// 					words[j] = Capitalize(words[j])
// 				}
// 			}

// 			// Удаляем команду "(cap, <число>)" из среза
// 			words = append(words[:i], words[i+2:]...)
// 			i-- // Уменьшаем индекс, так как мы удалили два элемента
// 		}
// 	}

// 	if words[i+1] == "(up)" {
// 		words[i] = ToUpperWord(words[i])
// 		words = append(words[:i+1], words[i+2:]...)
// 		i--
// 	}
// 	if strings.HasPrefix(words[i], "(up,") {
// 		// Извлекаем число из следующего элемента
// 		if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
// 			// Убираем ")" из числа
// 			numStr := strings.TrimSuffix(words[i+1], ")")
// 			num, err := strconv.Atoi(numStr)
// 			if err != nil {
// 				// Если число не распознано, пропускаем
// 				continue
// 			}

// 			// Проверяем, что количество слов для обработки не превышает доступных
// 			if i-num >= 0 {
// 				// Переписываем указанное количество слов с заглавной буквы
// 				for j := i - num; j < i; j++ {
// 					words[j] = ToUpperWord(words[j])
// 				}
// 			}

// 			// Удаляем команду "(cap, <число>)" из среза
// 			words = append(words[:i], words[i+2:]...)
// 			i-- // Уменьшаем индекс, так как мы удалили два элемента
// 		}
// 	}
// 	if words[i+1] == "(low)" {
// 		words[i] = ToLowerWord(words[i])
// 	}
// 	if words[i] == "(low)" {
// 		words[i] = strings.TrimPrefix(words[i], "(low)")
// 	}

// 	if words[i+1] == "(hex)" {
// 		// Преобразуем текущий элемент в десятичное число из шестнадцатеричной строки
// 		decimalValue, err := strconv.ParseInt(words[i], 16, 64)
// 		if err != nil {
// 			// Если произошла ошибка, выводим ее
// 			// fmt.Println("Ошибка при преобразовании:", err)
// 			continue
// 		}

// 		// Преобразуем число обратно в строку и записываем в words[i]
// 		words[i] = strconv.FormatInt(decimalValue, 10)
// 		// Удаляем элемент "(hex)" из массива
// 		words = append(words[:i+1], words[i+2:]...)
// 		i-- // уменьшаем индекс, чтобы снова проверить на текущей позиции
// 	}

// 	if words[i+1] == "(bin)" {
// 		decimalValue, err := strconv.ParseInt(words[i], 2, 64)
// 		if err != nil {
// 			// fmt.Println("Ошибка при преобразовании:", err)
// 			continue
// 		}
// 		words[i] = strconv.FormatInt(decimalValue, 10)
// 		words = append(words[:i+1], words[i+2:]...)
// 		i--
// 	}
// }

// // Собираем текст обратно
// editedText := strings.Join(words, " ")

// // Записываем в новый файл
// err = os.WriteFile(outputFile, []byte(editedText), 0o644)
// if err != nil {
// 	fmt.Println("Ошибка при записи файла:", err)
// 	return
// }

// fmt.Println("Файл успешно обработан и сохранён как", outputFile)
//}
