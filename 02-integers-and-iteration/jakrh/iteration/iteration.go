package iteration

import "strings"

const repeatCount = 5

func RepeatSlow(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func RepeatFast(character string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func RepeatWithCount(character string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(character)
	}
	return repeated.String()
}