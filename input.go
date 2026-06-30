package main

import (
	"bufio"
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

func FastScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	l := len(data)
	for i := 0; i < l; i++ {
		b := data[i]
		if b == ' ' || b == '\n' {
			return i + 1, data[0:i], nil
		}
	}
	if atEOF && l > 0 {
		return l, data[0:], nil
	}
	return 0, nil, nil
}
