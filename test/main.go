package main

import (
	"fmt"
	"strings"
)

func UpCommand(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(up)" {
			continue
		}

		if i+1 < len(words) && words[i+1] == "(up)" && words[i+1] != "(cap)" && words[i+1] != "(low)" && words[i+1] != "(bin)" && words[i+1] != "(hex)" {
			result = append(result, strings.ToUpper(word))
			i++
		} else {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	line := []string{"files", "(up)", "(up)", "(up)", "(low)", "(up)"}
	result := UpCommand(line)
	fmt.Println(result)

}
