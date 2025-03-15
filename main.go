package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func capitalizeWord(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	return strings.ToUpper(string(word[0])) + word[1:]
// }

// func ToUpperWord(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	return strings.ToUpper(word)
// }

// func ToLowerWord(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	return strings.ToLower(word)
// }

func main() {
	// 	inputFile := "sample.txt"
	// 	outputFile := "result.txt"

	// 	content, err := os.ReadFile(inputFile)
	// 	if err != nil {
	// 		fmt.Println("Ошибка при чтении файла:", err)
	// 		return
	// 	}
	// 	text := string(content)

	// 	// Разбиваем текст на слова
	// 	words := strings.Split(text, " ")

	// 	// Проходим по словам и ищем (cap)
	// 	for i := 0; i < len(words)-1; i++ {
	// 		if words[i+1] == "(cap)" {
	// 			words[i] = capitalizeWord(words[i])
	// 		}
	// 		if words[i] == "(cap)" {
	// 			words[i] = strings.TrimPrefix(words[i], "(cap)")
	// 		}
	// 		if words[i+1] == "(up)" {
	// 			words[i] = ToUpperWord(words[i])
	// 			words = append(words[:i+1], words[i+2:]...)
	// 			i--
	// 		}
	// 		if words[i+1] == "(low)" {
	// 			words[i] = ToLowerWord(words[i])
	// 		}
	// 		if words[i] == "(low)" {
	// 			words[i] = strings.TrimPrefix(words[i], "(low)")
	// 		}
	// 		if words[i+1] == "(hex)" {
	// 			// Преобразуем текущий элемент в десятичное число из шестнадцатеричной строки
	// 			decimalValue, err := strconv.ParseInt(words[i], 16, 64)
	// 			if err != nil {
	// 				// Если произошла ошибка, выводим ее
	// 				// fmt.Println("Ошибка при преобразовании:", err)
	// 				continue
	// 			}

	// 			// Преобразуем число обратно в строку и записываем в words[i]
	// 			words[i] = strconv.FormatInt(decimalValue, 10)
	// 			// Удаляем элемент "(hex)" из массива
	// 			words = append(words[:i+1], words[i+2:]...)
	// 			i-- // уменьшаем индекс, чтобы снова проверить на текущей позиции
	// 		}
	// 		if words[i+1] == "(bin)" {
	// 			decimalValue, err := strconv.ParseInt(words[i], 2, 64)
	// 			if err != nil {
	// 				// fmt.Println("Ошибка при преобразовании:", err)
	// 				continue
	// 			}
	// 			words[i] = strconv.FormatInt(decimalValue, 10)
	// 			words = append(words[:i+1], words[i+2:]...)
	// 			i--
	// 		}

	// 	}

	// 	// Собираем текст обратно
	// 	editedText := strings.Join(words, " ")

	// 	// Записываем в новый файл
	// 	err = os.WriteFile(outputFile, []byte(editedText), 0o644)
	// 	if err != nil {
	// 		fmt.Println("Ошибка при записи файла:", err)
	// 		return
	// 	}

	// fmt.Println("Файл успешно обработан и сохранён как", outputFile)
}
