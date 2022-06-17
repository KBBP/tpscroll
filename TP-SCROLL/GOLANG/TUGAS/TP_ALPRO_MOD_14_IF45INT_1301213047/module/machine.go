package main

import . "fmt"

var (
	ribbon string
	CC     byte
	EOP    bool
	index  int
)

func START() {
	index = 0
	CC = ribbon[index]
	EOP = string(CC) == "."
}

func ADV() {
	index++
	CC = ribbon[index]
	EOP = string(CC) == "."
}

func DIGIT() bool {
	return (string(CC) == "0") || (string(CC) == "1") || (string(CC) == "2") || (string(CC) == "3") || (string(CC) == "4") || (string(CC) == "5") || (string(CC) == "6") || (string(CC) == "7") || (string(CC) == "8") || (string(CC) == "9")
}

func main() {
	var num int
	var faculty string
	var gen, class, facState bool

	Scan(&ribbon)

	START()
	for !(EOP) {
		if !DIGIT() {
			faculty += string(CC)
			if faculty == "IF" {
				facState = true
			}
		}
		if string(CC) == "-" {
			ADV()
			for !(string(CC) == "-") && !(EOP) {
				if DIGIT() && !gen {
					gen = true
				} else if DIGIT() && num > 0 {
					class = true
				}
				ADV()
			}
			num++
			if num > 2 {
				class = false
			}
		}
		if !(string(CC) == "-") && !(EOP) {
			ADV()
		}
	}
	Println(facState && gen && class)
}
