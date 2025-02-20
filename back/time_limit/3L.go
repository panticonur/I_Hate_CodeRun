// https://coderun.yandex.ru/selections/backend/problems/gcd-and-lcm-yandex
// https://younglinux.info/algorithm/euclidean
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ScanInt(scanner *bufio.Scanner) int {
	return StrToInt(ScanWord(scanner))
}

func ScanWord(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	return scanner.Text()
}

func StrToInt(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func main3() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	x, y := ScanInt(scan), ScanInt(scan)
	log.Println("IN:", x, y)
	log.Println("NOD", NOD(x, y))
	log.Println("NOK", NOK(x, y))

	if x%y != 0 && y%x != 0 {
		fmt.Println(0)
		return
	}
	if x > y {
		x, y = y, x
	}

	d, k, q := NOD(x, y), NOK(x, y), 0
	for a := x; a <= y; a = a + x {
		log.Println(a)
		for b := y; b >= x; b = b - x {
			if NOD(a, b) == d && NOK(a, b) == k {
				q++
				log.Println(q, ")", a, b)
			}
		}
	}

	fmt.Println(q)
}

func NOD(x, y int) int {
	a, b := x, y
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func NOK(x, y int) int {
	a, b := x, y
	for a != b {
		if a > b {
			b = b + y
		} else {
			a = a + x
		}
	}
	return a
}
