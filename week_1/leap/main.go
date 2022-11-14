// This is a "stub" file. Це є маленький запуск на вашій згоді.
// It's not a complete solution though; You have to write some code.

// Package leap should have package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear повинен мати коментар documenting it.
func IsLeapYear(year int) bool {
	// Тут має бути рішення
	// написавши код - необхідно запустити тести
	// Ці коментарі можна видаляти
	// !ВАЖЛИВО - не забудьте виправити return

	return (year%400 == 0 || year%4 == 0 && year%100 != 0)
}
