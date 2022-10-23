package pangram

import (
	"strings"
)

// IsPangram checks if given string is a pangram
func IsPangram(s string) bool {
	// Тут має бути рішення
	// написавши код - необхідно запустити тести
	// Ці коментарі можна видаляти
	// !ВАЖЛИВО - не забудьте виправити return

	// Англійський алфавіт (довжина 26 символів)
	alphabet := string("abcdefghijklmnopqrstuvwxyz")
	// Кількість надходжень співпадінь при порівнянні
	accumulator := 0

	/* Якщо символ в адфавіті співпадає з символом в строці-аргументі,
	то сбільшуємо інкрементуємо accumulator
	*/
	for i := 0; i < len(alphabet); i++ {
		for j := 0; j < len(s); j++ {
			if string(alphabet[i]) == strings.ToLower(string(s[j])) {
				accumulator++
				break
			}
		}
	}
	// true = pangrama
	return (accumulator == len(alphabet))
}
