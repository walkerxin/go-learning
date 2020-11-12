package iteration

import "strings"

// Repeat the taked character return five
func Repeat(character string, frequency int) string {
	var repeated string
	for i:=0; i<frequency; i++ {
		repeated += character
	}
	return repeated
}

func ApiRepeat(character string, count int) string {
	return strings.Repeat(character, count)
}