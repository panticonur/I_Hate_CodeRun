// https://coderun.yandex.ru/selections/2024-summer-backend/problems/nearest-bus-stop
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

func FastScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	l := len(data)
	for i := 0; i < l; i++ {
		b := data[i]
		if b == ' ' || b == '\n' {
			return i + 1, data[0:i], nil
		}
	}
	if atEOF && l > 0 {
		return l, data[0:], nil
	}
	return 0, nil, nil
}

var scanner = func() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 0, 1048576)
	scanner.Buffer(buffer, 1048576)
	scanner.Split(bufio.ScanWords)
	return scanner
}()

func main() {
	n, k := ScanInt(), ScanInt()

	var S []int
	for i := 0; i < n; i++ {
		S = append(S, ScanInt())
	}

	for i := 0; i < k; i++ {
		z := ScanInt()

		l := 0
		for ; l < n; l++ {
			if S[l] > z {
				l = l - 1
				break
			}
		}

		r := n - 1
		for ; r >= 0; r-- {
			if S[r] < z {
				r = r + 1
				break
			}
		}

		// log.Println(l, r)
		if l == -1 && r == -1 {
			fmt.Println(1)
		} else if l == n && r == n {
			fmt.Println(n)
		} else if r == -1 {
			fmt.Println(l + 1)
		} else if l == n {
			fmt.Println(r + 1)
		} else {
			fmt.Println(l + 1)
		}
	}
}
