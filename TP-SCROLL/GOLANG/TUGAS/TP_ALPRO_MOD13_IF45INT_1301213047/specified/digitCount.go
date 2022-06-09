package main

import . "fmt"

var counter = 0

func main() {
	num, result := 0, 0

	Printf("Enter number: ")
	Scan(&num)
	result = countDig(num)
	Print("The digits: ", result)
}

func countDig(n int) int {
	if n > 0 {
		counter++
		countDig(n / 10)
	}

	return counter
}
