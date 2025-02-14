// https://coderun.yandex.ru/problem/trading-ya-intern
package main

import (
	"fmt"
	"sort"
)

func main324() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)
	// var words []string
	// for scanner.Scan() {
	//     words = append(words, scanner.Text())
	// }

	// var ints []int
	// for _, s :=  range words {
	//     i, err := strconv.Atoi(s)
	//     if err != nil {
	//         panic(err)
	//     }
	//     ints = append(ints, i)
	// }

	// n := ints[0]
	// // n, m := ints[0], ints[1]
	// a := ints[2:2+n]
	// b := ints[2+n:]

	a := []int{5, 10, 8, 4, 7, 2}
	b := []int{3, 1, 11, 18, 9}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	fmt.Println(a)
	fmt.Println(b)
	x := 0
	for i, bb := range b {
		if i >= len(a) {
			break
		}
		g := bb - a[i]
		if g <= 0 {
			break
		}
		x += g
		fmt.Println(x)
	}
	fmt.Println(x)
}
