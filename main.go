// https://coderun.yandex.ru/selections/backend/problems/diversity-scoring
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
	n := ScanInt()
	P := make(map[int]int, n)
	C := make([]int, n)
	for i := 1; i <= n; i++ {
		P[ScanInt()] = ScanInt()
	}

	M := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		c := P[ScanInt()]
		C[i] = c
		M[c] = false
	}

	log.Println(C)
	min := len(C)

	for i := 0; i < n; i++ {
		c := C[i]
		if M[c] {
			continue
		}

		for j, J := i+1, i; j < n; j++ {
			if C[j] == c {
				if j-J < min {
					min = j - J
				}
				J = j
			}
		}
		for j, J := i-1, i; j >= 0; j-- {
			if C[j] == c {
				if J-j < min {
					min = J - j
				}
				J = j
			}
		}

		M[c] = true
		if min <= 1 {
			break
		}
	}

	fmt.Println(min)
}
