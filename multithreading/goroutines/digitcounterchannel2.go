package main

import (
	"strings"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc func() string

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsWithChannel2(next nextFunc) counter {
	counted := make(chan pair)

	// начало решения

	go func() {
		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := next()
			digits := countDigits(word)
			counted <- pair{word: word, count: digits}
			if word == "" {
				break
			}
		}
	}()

	// Считайте значения из канала counted
	// и заполните stats.
	stats := counter{}
	for {
		count := <-counted
		if count.word == "" {
			break
		}
		stats[count.word] = count.count
	}
	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}

// wordGenerator возвращает генератор, который выдает слова из фразы
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsWithChannel2(next)
	printStats(stats)
}
