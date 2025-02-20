// https://coderun.yandex.ru/selections/backend/problems/mobilization
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
	n := ScanInt()
	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = ScanInt()
	}
	for i := 0; i < n; i++ {
		B[i] = ScanInt()
	}

	const TYPE_1, TYPE_2 = 1, 2
	T := make([]int, n)
	for sert, m := 1, ScanInt(); sert <= m; sert++ {
		{
			num := ScanInt()
			if ScanInt() == TYPE_1 {
				A[num-1] = A[num-1] + ScanInt()
			} else {
				B[num-1] = B[num-1] + ScanInt()
			}
		}

		// if m == 100 {
		// 	panic(m)
		// }

		max := 0
		for k := uint64(0); k < 1<<n; k++ {
			h := 0
			for i, r := 0, uint64(1); i < n; i, r = i+1, r<<1 {
				x := k & r
				x = x >> i
				if x == TYPE_1 {
					h++
				}
				T[i] = TYPE_1 + int(x)
				// log.Println(T, k, i, r, k&r, x)
			}
			if h != n/2 {
				continue
			}

			a, b := 0, 0
			for i, t := range T {
				if t == TYPE_1 {
					a += A[i]
				}
				if t == TYPE_2 {
					b += B[i]
				}
			}

			log.Println(T, a, b, a+b)
			if a+b > max {
				max = a + b
			}
		}

		fmt.Println(max)
	}
}
