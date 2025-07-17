// https://coderun.yandex.ru/selections/backend/problems/sorting-reverse-order
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

func main() {
	var host string
	var port int
	var a, b int
	fmt.Scan(&host, &port, &a, &b)

	url := fmt.Sprintf("%s:%d?a=%d&b=%d", host, port, a, b)
	responce, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := io.ReadAll(responce.Body)

	var values []int
	json.Unmarshal(data, &values)

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for _, n := range values {
		if n > 0 {
			fmt.Print(n)
			fmt.Print("\n")
		}
	}

}
