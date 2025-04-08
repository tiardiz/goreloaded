package firstproject

import (
	"regexp"
	"strings"
)

func MergeStrings(slice []string) []string {
	var result []string
	reStart := regexp.MustCompile(`\((cap|up|low)`)
	reEnd := regexp.MustCompile(`(\d+)\)`)

	i := 0
	for i < len(slice) {
		if reStart.MatchString(slice[i]) {
			if i+1 < len(slice) && reEnd.MatchString(slice[i+1]) {
				result = append(result, slice[i]+" "+slice[i+1])
				i++
			} else {
				result = append(result, slice[i])
			}
		} else {
			result = append(result, slice[i])
		}
		i++
	}
	result = SplitAfterParenthesis(result)
	return result

}

func SplitAfterParenthesis(input []string) []string {
	var result []string

	for _, item := range input {
		// Находим индекс символа ')'
		index := strings.Index(item, ")")
		if index != -1 {
			// Добавляем часть строки до ')'
			result = append(result, item[:index+1])

			// Добавляем часть строки после ')'
			if index+1 < len(item) {
				result = append(result, item[index+1:])
			}
		} else {
			result = append(result, item)
		}
	}

	return result
}
