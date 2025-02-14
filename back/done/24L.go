// https://coderun.yandex.ru/selections/backend/problems/server-error
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scan24(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		log.Fatal("scanner")
	}
	word := scanner.Text()
	x, err := strconv.Atoi(word)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func main24() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := scan24(scanner)
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = scan24(scanner)
		b[i] = scan24(scanner)
	}

	// n := 2
	// a := []int{50, 50}
	// b := []int{1, 2}

	// n := 3
	// a := []int{10, 30, 60}
	// b := []int{100, 10, 2}

	c := 0
	for i := 0; i < n; i++ {
		c += a[i] * b[i]
	}

	for i := 0; i < n; i++ {
		a, b := a[i], b[i]
		// fmt.Println(a, b)
		fmt.Println(float64(a*b) / float64(c))
	}
}
