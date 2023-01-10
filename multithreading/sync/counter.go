package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	count map[string]int
	sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.Lock()
	defer c.Unlock()
	c.count[str] += 1
}

func (c *Counter) Value(str string) int {
	c.Lock()
	defer c.Unlock()
	return c.count[str]
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.Lock()
	defer c.Unlock()
	for s, i := range c.count {
		fn(s, i)
	}
}

func NewCounter() *Counter {
	return &Counter{
		count: map[string]int{},
		Mutex: sync.Mutex{},
	}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
