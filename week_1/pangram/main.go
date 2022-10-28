package pangram

import (
	"strings"
)

/* 2 Additional solutions
1. Compare unicode
2. Use packet string
*/

// Solution 1
/*
func IsPangram(s string) bool {

	var (
		alphabet    = []rune("abcdefghijklmnopqrstuvwxyz")
		str         = strings.ToLower(string(s))
		accumulator = 0
	)

	for i := 0; i < len([]rune(alphabet)); i++ {
		for j := 0; j < len([]rune(str)); j++ {
			if []rune(alphabet)[i] == []rune(str)[j] {
				accumulator++
				break
			}
		}
	}

	return (accumulator == len(alphabet))
}
*/

// Solution 2
func IsPangram(s string) bool {
	var (
		alphabet  = []rune("abcdefghijklmnopqrstuvwxyz")
		str       = strings.ToLower(string(s))
		isPangram = true
	)

	for i := 0; i < len([]rune(alphabet)); i++ {
		if !strings.ContainsRune(string(str), alphabet[i]) {
			isPangram = false
			break
		}

	}
	return isPangram
}
