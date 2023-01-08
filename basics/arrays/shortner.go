package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	var res string

	if len(text) <= width {
		res = text
	} else {
		runes := []rune(text)
		runes = runes[:width]
		res = string(runes) + "..."
	}

	fmt.Println(res)
}
