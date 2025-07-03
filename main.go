// https://coderun.yandex.ru/selections/backend/problems/plane-boarding/
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
	buffer := make([]byte, 0, 1048576)
	scanner.Buffer(buffer, 1048576)
	scanner.Split(bufio.ScanWords)
	return scanner
}()

const (
	SEAT_A = 'A'
	SEAT_B = 'B'
	SEAT_C = 'C'
	SEAT_D = 'D'
	SEAT_E = 'E'
	SEAT_F = 'F'
	ROWS   = 30
)

type Passenger struct {
	Row     int
	Seat    rune
	A       int
	HallRow int
	Done    bool
	RowSeat string
}

func main() {
	count := ScanInt()
	passengers := make([]Passenger, count)

	for i := range count {
		passengers[i].A = ScanInt()

		rowSeat := ScanWord()
		rsLen := len(rowSeat)
		passengers[i].RowSeat = rowSeat

		seat := rowSeat[rsLen-1]
		passengers[i].Seat = rune(seat)

		row := rowSeat[:rsLen-1]
		passengers[i].Row = StrToInt(row)
	}

	log.Println(passengers)

	Hall := make([]*Passenger, ROWS+1) // коридор

	t := 0
	for {
		done := true
		for _, p := range passengers {
			done = done && p.Done
		}
		if done {
			fmt.Println(t)
			return
		}

		maxTime := 0
		for p := range passengers {
			passenger := &passengers[p]
			if passenger.Done {
				Hall[passenger.HallRow] = nil
				passenger.HallRow = 0
				continue
			}
			if passenger.HallRow < passenger.Row { // нужно двигаться
				if Hall[passenger.HallRow+1] == nil { // впереди есть куда шагнуть
					Hall[passenger.HallRow] = nil
					passenger.HallRow += 1
					Hall[passenger.HallRow] = passenger
					if maxTime < 1 {
						maxTime = 1
					}
				}
			} else if passenger.HallRow == passenger.Row { // садится
				passenger.Done = true
				//Hall[passenger.HallRow] = nil

				additionalTime := 0
				for n := range passengers { // ищем соседей по ряду
					neighbor := &passengers[n]
					if neighbor.Row == passenger.Row && neighbor.Done && neighbor != passenger {
						// отфильтрованы все седящие в ряду
						if passenger.Seat == SEAT_A || passenger.Seat == SEAT_B || passenger.Seat == SEAT_C {
							if neighbor.Seat == SEAT_A || neighbor.Seat == SEAT_B || neighbor.Seat == SEAT_C {
								if passenger.Seat < neighbor.Seat {
									if additionalTime == 0 {
										additionalTime = 5
									} else {
										additionalTime = 15
									}
								}
							}
						}
						if passenger.Seat == SEAT_D || passenger.Seat == SEAT_E || passenger.Seat == SEAT_F {
							if neighbor.Seat == SEAT_D || neighbor.Seat == SEAT_E || neighbor.Seat == SEAT_F {
								if passenger.Seat > neighbor.Seat {
									if additionalTime == 0 {
										additionalTime = 5
									} else {
										additionalTime = 15
									}
								}
							}
						}
					}
				}

				if maxTime < passenger.A+additionalTime {
					maxTime = passenger.A + additionalTime
				}
			}
		}

		t += maxTime
	}

}
