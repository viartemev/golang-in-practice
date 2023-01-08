package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения
	res := make(chan struct {
		int
		any
	}, len(funcs))
	result := make([]any, len(funcs))
	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его
	for i, f := range funcs {
		go func(i int, f func() any) {
			res <- struct {
				int
				any
			}{i, f()}
		}(i, f)
	}
	for {
		if len(res) == len(funcs) {
			close(res)
			break
		}
	}
	for r := range res {
		result[r.int] = r.any
	}
	return result
	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
