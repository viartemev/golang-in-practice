package main

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsWithChannel3(next nextFunc) counter {
	pending := make(chan string)
	counted := make(chan pair)

	// начало решения
	stats := counter{}
	// отправляет слова на подсчет
	go func() {
		// Пройдите по словам и отправьте их
		// в канал pending
		for {
			word := next()
			pending <- word
			if word == "" {
				break
			}
		}
	}()

	// считает цифры в словах
	go func() {
		// Считайте слова из канала pending,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := <-pending
			digits := countDigits(word)
			counted <- pair{word: word, count: digits}
			if word == "" {
				break
			}
		}
	}()

	// Считайте значения из канала counted
	// и заполните stats.
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

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsWithChannel3(next)
	printStats(stats)
}
