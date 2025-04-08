package main

import (
	"fmt"
	"os"

	firstproject "firstproject/funcs"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	text := string(content)

	processedText := firstproject.ProcessText(text)

	err = os.WriteFile(outputFile, []byte(processedText), 0o644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Обработка завершена. Результат сохранен в", outputFile)
}
