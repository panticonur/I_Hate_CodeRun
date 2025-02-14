// https://coderun.yandex.ru/selections/backend/problems/pairwise-xor
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scan7(scanner *bufio.Scanner) int {
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

func scan7f() int {
	x := 0
	_, err := fmt.Scan(&x)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func main7() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	T := scan7(scanner)

	for t := 0; t < T; t++ {
		n := scan7(scanner)
		a := make([]int, n)

		for i := 0; i < n; i++ {
			a[i] = scan7(scanner)
		}

		// ints := []int{1, 5, 1, 2, 4, 8, 1}
		// fmt.Println(n)
		// fmt.Println(a)

		m := 0xFFFFFFFF
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					x := a[i] ^ a[j]
					if x < m {
						m = x
					}
				}
			}
			if m == 0 {
				break
			}
		}
		fmt.Println(m)
	}
}
