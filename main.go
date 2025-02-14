// https://coderun.yandex.ru/selections/backend/problems/calendar-formatting
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	x, n := 1, ScanInt()
	w := ScanWord()
	if w == "Tuesday" {
		x = 0
	} else if w == "Wednesday" {
		x = -1
	} else if w == "Thursday" {
		x = -2
	} else if w == "Friday" {
		x = -3
	} else if w == "Saturday" {
		x = -4
	} else if w == "Sunday" {
		x = -5
	}
	for x <= n {
		var b strings.Builder
		for i := 1; i <= 7; i++ {
			s := strconv.Itoa(x)
			if x > 0 && x < 10 {
				b.WriteByte('.')
			}
			if x > 0 && x <= n {
				b.WriteString(s)
			} else {
				if x <= n {
					b.WriteString("..")
				} else {
					b.WriteString("  ")
				}
			}
			if i < 7 {
				b.WriteByte(' ')
			}
			x++
		}
		fmt.Println(b.String())
	}
}
