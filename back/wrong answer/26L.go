// https://coderun.yandex.ru/selections/backend/problems/symbols-set-min-susbstr
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scan26(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		log.Fatal("scanner")
	}
	return scanner.Text()
}

func main26() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	S := scan26(scanner)
	C := scan26(scanner)

	// S := "abba"
	// C := "ab"

	// S := "accb"
	// C := "cab"

	X := len(S)
	F := false

	for i := 0; i <= len(S); i++ {
		M := make(map[byte]bool, len(C))
		for m := 0; m < len(C); m++ {
			M[C[m]] = false
		}
		x := 0
		for j := i; j < len(S); j++ {
			s := S[j]
			if _, b := M[s]; b {
				M[s] = true
				x++
			} else {
				break
			}
		}
		// fmt.Println(x)
		full := true
		for _, f := range M {
			if !f {
				full = false
			}
		}
		if full && x > 0 {
			F = true
			if x < X {
				X = x
			}
		}
	}
	// fmt.Println()
	if F {
		fmt.Println(X)
	} else {
		fmt.Println(0)
	}
}
