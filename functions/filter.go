package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	result := make([]int, 0)
	for _, i := range iterable {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	src := readInput()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	res := filter(func(i int) bool { return i%2 == 0 }, src)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
