// https://coderun.yandex.ru/selections/backend/problems/median-out-of-three
package main

import "fmt"

func main6() {
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

	// a, b, c := 1, 2, 3
	// a, b, c := 3, 2, 1
	// a, b, c := 2, 3, 1
	// a, b, c := 2, 1, 3
	// a, b, c := 1, 3, 2
	// a, b, c := 3, 1, 2

	// a, b, c := 3, 2, 3
	// a, b, c := 3, 2, 3
	// a, b, c := 2, 3, 3
	// a, b, c := 2, 3, 3
	// a, b, c := 3, 3, 2
	// a, b, c := 3, 3, 2

	// fmt.Println(a, b, c)

	ints := []int{1, 3, 2}

	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			if ints[i] > ints[j] {
				// fmt.Println(ints)
				ints[i], ints[j] = ints[j], ints[i]
				// fmt.Println(ints)
				// fmt.Println()
			}
		}
	}

	// fmt.Println()
	// fmt.Println(ints)
	fmt.Println(ints[1])
}
