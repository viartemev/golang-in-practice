package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	cancel := make(chan struct{})
	ticker := time.NewTicker(dur)

	cancelFn := func() {
		select {
		case <-cancel:
		default:
			close(cancel)
			ticker.Stop()
		}
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			case <-cancel:
				return
			}
		}
	}()

	return cancelFn
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}
