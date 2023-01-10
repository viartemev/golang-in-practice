package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	ticker := time.NewTicker(time.Duration(1000/limit) * time.Millisecond)
	cancelChannel := make(chan struct{})
	cancel = func() {
		select {
		case <-cancelChannel:
		default:
			close(cancelChannel)
			ticker.Stop()
		}
	}
	handle = func() error {
		select {
		case <-ticker.C:
			go fn()
			return nil
		case <-cancelChannel:
			return ErrCanceled
		}
	}
	return
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := withRateLimit(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
