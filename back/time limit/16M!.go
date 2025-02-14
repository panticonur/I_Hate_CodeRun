package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

// https://coderun.yandex.ru/selections/backend/problems/find-rle-string-length

func scanWord_16(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	return scanner.Text()
}

func integer_16(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func scanWords_16(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func main16() {
	bufMax := 1_048_576
	scannerBuf := make([]byte, 0, bufMax)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(scannerBuf, bufMax)
	scanner.Split(scanWords_16)

	rle := scanWord_16(scanner)
	Q := integer_16(scanWord_16(scanner))
	// rle := "a2bc3a"
	// Q := 5

	// if len(rle) != 100_984 && len(rle)==100_011 { // 23 24
	// 	log.Fatal(0)
	// }

	B := make([]int, 0, 65_536)
	E := make([]int, 0, 65_536)
	K := make([]int, 0, 65_536)
	{
		var buf bytes.Buffer
		b := 0
		for i, t, l := 0, 0, len(rle); i < l; i++ {
			c := rle[i]
			if c >= '0' && '9' >= c {
				buf.WriteByte(c)
				if b == 0 {
					t++
					b = t
				}
			} else {
				if b == 0 {
					t++
					B = append(B, t)
					E = append(E, t)
					K = append(K, 1)

				}
				bl := buf.Len()
				if bl > 0 {
					n, _ := strconv.Atoi(buf.String())
					buf.Reset()
					t += n - 1
					B = append(B, b)
					E = append(E, t)
					K = append(K, bl+1)
					b = 0
				}
			}
		}
	}
	lenS := len(B)

	for q := 0; q < Q; q++ {
		l, _ := strconv.Atoi(scanWord_16(scanner))
		r, _ := strconv.Atoi(scanWord_16(scanner))
		/* l, r := 0, 0
		if q == 0 {
			l, r = 1, 7
		} else if q == 1 {
			l, r = 5, 7
		} else if q == 2 {
			l, r = 1, 2
		} else if q == 3 {
			l, r = 3, 5
		} else if q == 4 {
			l, r = 4, 4
		} */

		sl, j := 0, lenS
		for sl < j {
			h := (sl + j) >> 1
			if E[h] < l {
				sl = h + 1
			} else {
				j = h
			}
		}

		sr, j := sl, lenS
		for sr < j {
			h := (sr + j) >> 1
			if E[h] < r {
				sr = h + 1
			} else {
				j = h
			}
		}

		if sr == sl {
			X, x := 0, r-l+1
			if x > 1 {
				num := strconv.Itoa(x)
				X += len(num)
			}
			X++
			fmt.Println(X)
			continue
		}

		X, x := 0, E[sl]-l+1
		if x > 1 {
			num := strconv.Itoa(x)
			X += len(num)
		}
		X++

		x = r - B[sr] + 1
		if x > 1 {
			num := strconv.Itoa(x)
			X += len(num)
		}
		X++

		// for _, k := range K[sl+1 : sr] {
		// 	X += k
		// }

		for p, r, s := unsafe.Pointer(&K[sl+1]), unsafe.Pointer(&K[sr]),
			unsafe.Sizeof(X); p != r; p = unsafe.Add(p, s) {
			X += *(*int)(p)
		}

		fmt.Println(X)
	}
}
