// https://coderun.yandex.ru/problem/merge-jsons-2
package main

import (
	"bufio"
	"encoding/json"
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

type Offer struct {
	OfferID string `json:"offer_id"`
	SKU     int    `json:"market_sku"`
	Price   int    `json:"price"`
}

type Feed struct {
	Offers []Offer `json:"offers"`
}

func main() {
	n, m := 0, 0
	fmt.Sscanf(ScanLine(), "%d %d", &n, &m)
	// log.Println(n, m)
	var offers []Offer
	for i := 0; i < n; i++ {
		var feed Feed
		err := json.Unmarshal([]byte(ScanLine()), &feed)
		if err != nil {
			panic(err)
		}
		offers = append(offers, feed.Offers...)
		if len(offers) >= m {
			offers = offers[:m]
			out := Feed{
				Offers: offers,
			}
			str, err := json.Marshal(out)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(str))
			break
		}
	}
}
