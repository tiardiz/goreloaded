package firstproject

import (
	"os"
)

// Функция для чтения содержимого файла и возврата как строки
func ReadFile(filePath string) (string, error) {
	// Чтение содержимого файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err // Возвращаем ошибку, если не удается прочитать файл
	}

	// Преобразуем срез байтов в строку и возвращаем
	return string(data), nil
}

// func main() {
// 	// Проверка на количество аргументов
// 	if len(os.Args) < 2 {
// 		fmt.Println("Пожалуйста, укажите путь к файлу в качестве аргумента.")
// 		return
// 	}

// 	// Получаем путь к файлу из аргументов командной строки
// 	filePath := os.Args[1]

// 	// Вызываем функцию для чтения файла
// 	content, err := ReadFile(filePath)
// 	if err != nil {
// 		// Если произошла ошибка, выводим её
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}

// 	// Выводим содержимое файла
// 	fmt.Println(content)
// }
