package main

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsWithChannel4(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

// начало решения

// submitWords отправляет слова на подсчет
func submitWords(next nextFunc, pending chan<- string) {
	for {
		word := next()
		pending <- word
		if word == "" {
			break
		}
	}
}

// countWords считает цифры в словах
func countWords(pending <-chan string, counted chan<- pair) {
	for {
		word := <-pending
		digits := countDigits(word)
		counted <- pair{word: word, count: digits}
		if word == "" {
			break
		}
	}
}

// fillStats готовит итоговую статистику
func fillStats(counted <-chan pair) counter {
	stat := counter{}
	for {
		count := <-counted
		if count.word == "" {
			break
		}
		stat[count.word] = count.count
	}
	return stat
}

// конец решения

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsWithChannel4(next)
	printStats(stats)
}
