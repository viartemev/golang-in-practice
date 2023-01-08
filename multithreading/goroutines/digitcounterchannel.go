package main

import (
	"strings"
)

// countDigitsInWords считает количество цифр в словах фразы
func countDigitsInWordsWithChannel(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)

	// начало решения
	var stats = make(map[string]int)

	go func() {
		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for _, word := range words {
			digits := countDigits(word)
			counted <- digits
		}
	}()

	// Считайте значения из канала counted
	// и заполните stats.
	for _, word := range words {
		stats[word] = <-counted
	}

	// В результате stats должна содержать слова
	// и количество цифр в каждом.
	// конец решения

	return stats
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWordsWithChannel(phrase)
	printStats(stats)
}
