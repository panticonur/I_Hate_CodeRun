// https://coderun.yandex.ru/selections/backend/problems/exactly-one-occur
package main

import (
	"fmt"
)

func main2() {
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

	ints := []int{5, 1, 2, 3, 4, 4}

	n := ints[0]
	a := ints[1:]

	fmt.Println(n, a)

	m := map[int]int{}

	for _, i := range a {
		m[i] = 0
		//if v, ok:= m[i];
	}

	for _, i := range a {
		v := m[i]
		m[i] = v + 1
	}

	unic := 0
	for _, i := range a {
		v := m[i]
		if v == 1 {
			unic++
		}
	}

	fmt.Println(m)
	fmt.Println(unic)

}
