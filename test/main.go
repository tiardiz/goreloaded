package main

import (
	"fmt"
	"strings"
)

func main() {
	content := `keekeek, jxjj
keekeek, jxjj
she said. '
she said.'

hi ' hi ' hi
hi 'hi' hi`

	// Разделение на строки
	lines := strings.Split(content, "\n")

	// Вывод каждой строки
	for _, line := range lines {
		// fmt.Println(lines[i])
		fmt.Println(line)
	}
}
