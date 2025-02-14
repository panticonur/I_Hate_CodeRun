// https://coderun.yandex.ru/selections/backend/problems/a-1-find-most-frequent/description
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ScanWord() string {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	return scanner.Text()
}

func ScanInt() int {
	return StrToInt(ScanWord())
}

func StrToInt(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

var scanner = func() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return scanner
}()

func main() {
	M := make(map[int]int)
	max := 0
	X := 0
	for i, n := 0, ScanInt(); i < n; i++ {
		x := ScanInt()
		m := M[x] + 1
		if m > max {
			max = m
			X = x
		} else if m == max && x > X {
			X = x
		}
		M[x] = m
	}
	fmt.Println(X)
}
