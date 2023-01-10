package main

import (
	"context"
	"fmt"
	"strings"
	"unicode"
)

// информация о количестве цифр в каждом слове
type counter map[string]int

// слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// начало решения

// считает количество цифр в словах
func countDigitsInWords(ctx context.Context, words []string) counter {
	pending := submitWords(ctx, words)
	counted := countWords(ctx, pending)
	return fillStats(ctx, counted)
}

// отправляет слова на подсчет
func submitWords(ctx context.Context, words []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, word := range words {
			select {
			case out <- word:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// считает цифры в словах
func countWords(ctx context.Context, in <-chan string) <-chan pair {
	out := make(chan pair)
	go func() {
		defer close(out)
		for word := range in {
			select {
			case out <- pair{word, countDigits(word)}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// готовит итоговую статистику
func fillStats(ctx context.Context, in <-chan pair) counter {
	stats := counter{}
	for p := range in {
		select {
		case <-ctx.Done():
			return stats
		default:
			stats[p.word] = p.count
		}
	}
	return stats
}

// конец решения

// считает количество цифр в слове
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	words := strings.Fields(phrase)

	ctx := context.Background()
	stats := countDigitsInWords(ctx, words)
	fmt.Println(stats)
}
