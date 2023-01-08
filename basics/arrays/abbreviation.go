package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	unicode "unicode"
)

func main() {
	phrase := readString()
	var abbr []rune
	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...
	fields := strings.Fields(phrase)
	for _, field := range fields {
		runes := []rune(field)
		firstLetter := runes[0]
		if unicode.IsLetter(firstLetter) {
			abbr = append(abbr, unicode.ToUpper(firstLetter))
		}
	}

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
