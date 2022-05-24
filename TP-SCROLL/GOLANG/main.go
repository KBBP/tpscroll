package main

import "fmt"

func main() {
	var myArray [6]int
	myArray[1] = 2
	myArray[3] = 4
	myArray[5] = 7

	for i := 0; i <= 6; i++ {
		if myArray[i] < 5 {
			fmt.Println(myArray[i])
		}
	}
}
