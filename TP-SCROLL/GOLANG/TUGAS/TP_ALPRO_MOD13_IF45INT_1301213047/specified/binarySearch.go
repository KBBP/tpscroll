package main

import . "fmt"

func main() {
	var num int

	Println("Enter number: ")
	Scanln(&num)
	Println("Your Binary:", binaryConv(num))
}

func binaryConv(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n % 2) + 10*binaryConv(n/2)
	}
}
