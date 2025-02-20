// https://coderun.yandex.ru/selections/backend/problems/dsu-ya-intern/description
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
	n, m := ScanInt(), ScanInt()
	M := []*struct{ x, y int }{}
	for i := 0; i < m; i++ {
		M = append(M, &struct{ x, y int }{ScanInt() - 1, ScanInt() - 1})
	}
	q := ScanInt()
	Q := make([]int, q)
	for i := 0; i < q; i++ {
		Q[i] = ScanInt() - 1
	}

	for _, q := range Q {
		M[q] = nil

		L := make([][]int, n)
		for _, l := range M {
			if l != nil {
				L[l.x] = append(L[l.x], l.y)
				L[l.y] = append(L[l.y], l.x)
			}
		}
		log.Println(L)

		var mark func([]int, int, int)
		mark = func(A []int, a int, c int) {
			A[a] = c
			for _, l := range L[a] {
				if A[l] != c {
					mark(A, l, c)
				}
			}
		}

		A := make([]int, n)
		for i := 0; i < n; i++ {
			mark(A, i, i+1)
			log.Println(A)
		}

		min, max := n, 0
		for _, a := range A {
			if a < min {
				min = a
			}
			if a > max {
				max = a
			}
		}
		log.Println(max, min, max-min+1)
		fmt.Println(max - min + 1)
	}
}
