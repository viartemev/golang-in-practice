package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(duration time.Duration, fn func()) func() {
	canceled := make(chan struct{})

	timer := time.NewTimer(duration)
	go func() {
		select {
		case <-timer.C:
			fn()
		case <-canceled:
		}
	}()

	cancel := func() {
		if !timer.Stop() {
			return
		}
		close(canceled)
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
