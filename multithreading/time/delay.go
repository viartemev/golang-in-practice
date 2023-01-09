package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	canceled := make(chan struct{})

	go func() {
		time.Sleep(dur)
		select {
		case <-canceled:
			return
		default:
			fn()
		}
	}()

	cancel := func() {
		select {
		case <-canceled:
			return
		default:
			close(canceled)
		}
	}

	return cancel
}

// конец решения

func main() {
	rand.Seed(time.Now().Unix())

	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		cancel()
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
