// https://coderun.yandex.ru/selections/backend/problems/dayofweek-ya-intern
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
	buffer := make([]byte, 0, 1048576)
	scanner.Buffer(buffer, 1048576)
	scanner.Split(bufio.ScanWords)
	return scanner
}()

func main() {
	mm := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	md := map[int]string{
		0: "Saturday",
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Wednesday",
		5: "Thursday",
		6: "Friday",
	}

	for scanner.Scan() {
		day := StrToInt(scanner.Text())
		month := ScanWord()
		year := ScanInt()
		log.Println(day, month, year)

		q := day
		m := mm[month]
		if m < 3 {
			year -= 1
			m += 12
		}
		k := year % 100
		j := year / 100
		log.Println(k, j, m)

		h := (q + (13*(m+1))/5 + k + k/4 + j/4 - 2*j) % 7
		fmt.Println(md[h])
	}

}
