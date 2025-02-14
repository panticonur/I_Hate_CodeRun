// https://coderun.yandex.ru/selections/backend/problems/mew-http
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
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
	w1 := ScanWord()
	w2 := ScanWord()
	w3 := ScanWord()
	w4 := ScanWord()

	cli := &http.Client{}
	a1, a2 := f2(cli, w1, w2)
	a3 := f1(cli, w3)
	a4 := f1(cli, w4)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}

func f1(client *http.Client, w1 string) string {
	req, _ := http.NewRequest("MEW", "http://127.0.0.1:7777", nil)
	req.Header.Add("X-Cat-Variable", w1)
	resp, _ := client.Do(req)
	v := resp.Header.Values("X-Cat-Value")
	log.Println(w1, v)
	return v[0]
}

func f2(client *http.Client, w1, w2 string) (string, string) {
	req, _ := http.NewRequest("MEW", "http://127.0.0.1:7777", nil)
	req.Header.Add("X-Cat-Variable", w1)
	req.Header.Add("X-Cat-Variable", w2)
	resp, _ := client.Do(req)
	v := resp.Header.Values("X-Cat-Value")
	log.Println(w1, w2, v)
	return v[0], v[1]
}
