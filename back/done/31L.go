// https://coderun.yandex.ru/selections/backend/problems/checkers
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
	const white = 1
	const black = 2
	n, m := ScanInt(), ScanInt()
	D := make([][]int, n)
	for i := 0; i < n; i++ {
		D[i] = make([]int, m)
	}

	W := []struct{ n, m int }{}
	for i, w := 0, ScanInt(); i < w; i++ {
		n, m := ScanInt()-1, ScanInt()-1
		W = append(W, struct{ n, m int }{n, m})
		D[n][m] = white
	}
	B := []struct{ n, m int }{}
	for i, b := 0, ScanInt(); i < b; i++ {
		n, m := ScanInt()-1, ScanInt()-1
		B = append(B, struct{ n, m int }{n, m})
		D[n][m] = black
	}
	C := B
	e := white
	if ScanWord() == "white" {
		C = W
		e = black
	}

	for _, l := range D {
		log.Println(l)
	}
	log.Println()

	x := "No"
	for _, c := range C {
		x1, y1 := c.n-1, c.m+1
		x2, y2 := c.n-2, c.m+2
		if 0 <= x2 && y2 < m && D[x1][y1] == e && D[x2][y2] == 0 {
			x = "Yes"
			break
		}

		x1, y1 = c.n+1, c.m+1
		x2, y2 = c.n+2, c.m+2
		if x2 < n && y2 < m && D[x1][y1] == e && D[x2][y2] == 0 {
			x = "Yes"
			break
		}

		x1, y1 = c.n+1, c.m-1
		x2, y2 = c.n+2, c.m-2
		if x2 < n && 0 <= y2 && D[x1][y1] == e && D[x2][y2] == 0 {
			x = "Yes"
			break
		}

		x1, y1 = c.n-1, c.m-1
		x2, y2 = c.n-2, c.m-2
		if 0 <= x2 && 0 <= y2 && D[x1][y1] == e && D[x2][y2] == 0 {
			x = "Yes"
			break
		}
	}

	for _, l := range D {
		log.Println(l)
	}

	fmt.Println(x)
}
