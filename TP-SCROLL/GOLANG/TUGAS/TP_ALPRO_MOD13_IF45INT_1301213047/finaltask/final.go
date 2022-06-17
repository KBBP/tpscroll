package main

import "fmt"

var i = 2

func main() {
	var input int
	fmt.Scan(&input)
	harmonicLineV1(input, 2)
	harmonicLineV2(input)
}

func harmonicLineV1(n, i int) {
	if i == 2 {
		fmt.Printf("1 ")
	}

	if i <= n {
		fmt.Printf("+ 1/%v ", i)
		harmonicLineV1(n, i+1)
	}
}

func harmonicLineV2(n int) {
	if i == 2 {
		fmt.Printf("1 ")
	}

	if i <= n {
		fmt.Printf("+ 1/%v ", i)
		i += 1
		harmonicLineV2(n)
	}
}
