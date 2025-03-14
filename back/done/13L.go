// https://coderun.yandex.ru/selections/2024-summer-backend/problems/couple-of-letters
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
	var words []string
	for scanner.Scan() {
		w := scanner.Text()
		words = append(words, w)
	}
	log.Println(words)

	m := make(map[string]int)
	max := 0
	pair := ""
	for _, w := range words {
		for i, l := 0, len(w)-1; i < l; i++ {
			ab := w[i : i+2]
			p := m[ab] + 1
			m[ab] = p
			if p > max {
				max = p
				pair = ab
			}
			if p == max && ab > pair {
				pair = ab
			}
			log.Println(ab)
		}
	}
	log.Println(m)
	fmt.Println(pair)
}
