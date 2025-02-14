// https://coderun.yandex.ru/selections/backend/problems/meetings/description
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

type query struct {
	type_    string
	day      int
	hour     int
	mimute   int
	start    int
	stop     int
	duration int
	names    []string
}

func main() {
	var Q []*query
	M := make(map[string][]*query)
	for i, n := 0, ScanInt(); i < n; i++ {
		q := new(query)
		q.type_ = ScanWord()
		if q.type_ == "APPOINT" {
			q.day = ScanInt()
			time := strings.Split(ScanWord(), ":")
			q.hour = StrToInt(time[0])
			q.mimute = StrToInt(time[1])
			q.duration = ScanInt()
			q.start = q.day*24*60 + q.hour*60 + q.mimute
			q.stop = q.start + q.duration
			for j, k := 0, ScanInt(); j < k; j++ {
				q.names = append(q.names, ScanWord())
			}
		} else {
			q.day = ScanInt()
			q.names = append(q.names, ScanWord())
		}

		if q.type_ == "APPOINT" {
			X := make([]string, len(q.names))
			for _, p := range Q {
				if q.start < p.stop && q.stop > p.start {
					for i, x := range q.names {
						if P, ok := M[x]; ok {
							for _, p := range P {
								if q.start < p.stop && q.stop > p.start {
									X[i] = x
								}
							}
						}
					}
				}
			}
			ok := true
			for _, x := range X {
				if len(x) > 0 {
					ok = false
				}
			}
			if ok {
				fmt.Println("OK")
				for _, x := range q.names {
					M[x] = append(M[x], q)
				}
			} else {
				fmt.Println("FAIL")
				var buf strings.Builder
				for i, x := range X {
					if len(x) > 0 {
						buf.WriteString(x)
						if i < len(X)-1 {
							buf.WriteString(" ")
						}
					}
				}
				fmt.Println(buf.String())
			}

		} else {
			if P, ok := M[q.names[0]]; ok {
				start, stop := q.day*24*60, (q.day+1)*24*60-1
				var Q []*query
				for _, p := range P {
					if start < p.stop && stop > p.start {
						Q = append(Q, p)
					}
				}
				sort.Slice(Q, func(i, j int) bool { return Q[i].start < Q[j].start })
				for _, q := range Q {
					var buf strings.Builder
					buf.WriteString(
						fmt.Sprintf("%02d:%02d %d ", q.hour, q.mimute, q.duration))
					for i, y := range q.names {
						buf.WriteString(y)
						if i < len(q.names)-1 {
							buf.WriteString(" ")
						}
					}
					fmt.Println(buf.String())
				}
			}
		}

		Q = append(Q, q)
	}

}
