package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

// Функция обработки файла
func processFile(inputFile, outputFile string) {
	// Открываем входной файл
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаем выходной файл
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outFile.Close()

	// Чтение файла построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Применяем преобразования
		line = applyTransformations(line)

		// Записываем результат в выходной файл
		_, err := outFile.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

// Функция для применения преобразований
func applyTransformations(line string) string {
	// Преобразование "it" в "It", "IT" в "IT", и т.д.
	line = strings.ReplaceAll(line, "it", "It")
	line = strings.ReplaceAll(line, "IT", "IT")

	// Преобразование шестнадцатеричных и двоичных чисел в десятичные
	line = convertHexAndBinToDecimal(line)

	// Исправление пунктуации
	line = fixPunctuation(line)

	return line
}

// Функция для преобразования шестнадцатеричных и двоичных чисел
func convertHexAndBinToDecimal(line string) string {
	// Ищем все числа в формате (hex), (bin)
	re := regexp.MustCompile(`\((hex|bin)\s*([0-9a-fA-F]+)\)`)
	line = re.ReplaceAllStringFunc(line, func(match string) string {
		// Получаем тип числа и само число
		reMatch := regexp.MustCompile(`\((hex|bin)\s*([0-9a-fA-F]+)\)`).FindStringSubmatch(match)
		numType := reMatch[1]
		num := reMatch[2]

		var decimalValue int
		var err error
		if numType == "hex" {
			decimalValue, err = strconv.ParseInt(num, 16, 0)
		} else if numType == "bin" {
			decimalValue, err = strconv.ParseInt(num, 2, 0)
		}

		if err != nil {
			return match
		}

		return strconv.Itoa(int(decimalValue))
	})

	return line
}

// Функция для исправления пунктуации
func fixPunctuation(line string) string {
	// Убираем пробелы перед запятой и точками
	line = strings.ReplaceAll(line, " ,", ",")
	line = strings.ReplaceAll(line, " .", ".")
	return line
}

func main() {
	// Обрабатываем файл
	processFile("sample.txt", "result.txt")
}