// https://coderun.yandex.ru/selections/2024-summer-backend/problems/divisors-number
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ScanWord() string {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	return scanner.Text()
}

var scanner = func() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return scanner
}()

func main() {
	n := ScanInt()
	max, x := 0, 0
	for i := 1; i <= n; i++ {
		c := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				c++
			}
		}
		if c >= max {
			max = c
			x = i
		}
		log.Println(i, c)
	}
	fmt.Println(x)
	fmt.Println(max)
}
