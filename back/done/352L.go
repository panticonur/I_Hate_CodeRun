// https://coderun.yandex.ru/problem/rocks-and-jewels
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ScanLine() string {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	return scanner.Text()
}

var scanner = func() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	return scanner
}()

func main() {
	J, S := ScanLine(), ScanLine()
	count := 0
	for _, s := range S {
		m := make(map[rune]bool)
		for _, j := range J {
			if s == j {
				if _, exist := m[j]; !exist {
					m[j] = true
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
