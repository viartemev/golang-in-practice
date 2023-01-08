package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer fmt.Println("generate done")
		defer close(out)
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				break
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer fmt.Println("unique done")
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				if !ok {
					break
				}
				if unique(word) {
					out <- word
				}
			case <-cancel:
				break
			}
		}
	}()
	return out
}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}
	return true
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer fmt.Println("reverse done")
		defer close(out)
		for word := range in {
			select {
			case out <- reverseString(word):
			case <-cancel:
				break
			}
		}
	}()
	return out
}

func reverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}
	// return the reversed string.
	return string(rns)
}

// объединяет c1 и c2 в общий канал
func mergeWithCancellation(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		defer fmt.Println("merge done")
		for {
			select {
			case out <- <-c1:
			case out <- <-c2:
			case <-cancel:
				break
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for word := range in {
		if n > 0 {
			fmt.Println(word)
			n--
		} else {
			break
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := mergeWithCancellation(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
