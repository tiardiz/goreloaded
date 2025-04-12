package firstproject

import (
	"regexp"
	"strings"
)

func ControlCheck(text string) string {
	if strings.Contains(text, "(cap)") || strings.Contains(text, "(up)") || strings.Contains(text, "(low)") {
		return ProcessText(text)
	}

	re := regexp.MustCompile(`\((cap|up|low),\s*\d+\)`)
	if re.MatchString(text) {
		return ProcessText(text)
	}

	return text
}
