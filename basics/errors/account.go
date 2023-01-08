package main

// не меняйте импорты, они нужны для проверки
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var acc account
	var trans []int
	var err error
	acc, trans, err = parseInput()
	fmt.Print("-> ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(acc, trans)
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err1 := parseAccount(accSrc)
	if err1 != nil {
		return account{}, nil, err1
	}
	trans, err2 := parseTransactions(transSrc)
	if err2 != nil {
		return account{}, nil, err2
	}
	return acc, trans, nil
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return account{}, err1
	}
	overdraft, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		return account{}, err2
	}
	if overdraft < 0 {
		return account{}, errors.New("expect overdraft >= 0")
	}
	if balance < -overdraft {
		return account{}, errors.New("balance cannot exceed overdraft")
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return trans, err
		}
		trans[idx] = t
	}
	return trans, nil
}
