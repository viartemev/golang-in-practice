package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)
	var lang string
	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}
	fmt.Println(lang)
}
