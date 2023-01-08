package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say печатает фразу от имени обработчика
func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

// makePool создает пул на n обработчиков
// возвращает функции handle и wait
func makePool(n int, handler func(int, string)) (func(string), func()) {
	// начало решения
	pool := make(chan int, n)
	for i := 0; i < n; i++ {
		pool <- i
	}
	// создайте пул на n обработчиков
	// используйте для канала имя pool и тип chan int
	// определите функции handle() и wait()
	handle := func(str string) {
		id := <-pool
		go func() {
			handler(id, str)
			pool <- id
		}()
	}
	// handle() выбирает токен из пула
	// и обрабатывает переданную фразу через handler()

	// wait() дожидается, пока все токены вернутся в пул
	wait := func() {
		for i := 0; i < n; i++ {
			<-pool
		}
	}
	// конец решения

	return handle, wait
}

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	handle, wait := makePool(2, say)
	for _, phrase := range phrases {
		handle(phrase)
	}
	wait()
}
