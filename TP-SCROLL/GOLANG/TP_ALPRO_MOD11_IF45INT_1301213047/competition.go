package main

import (
	"fmt"
)

func main() {
	var sizeArray int

	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&sizeArray)

	var myArr = make([]int, sizeArray)

	for i := 0; i < sizeArray; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scan(&myArr[i])
	}

	// Finding Median Function

	for i := 0; i < sizeArray; i++ {
		for x := 1; x < len(myArr); i++ {

		}
	}

	// Array Template: [6, 3, 8, 2](4) = sizeOf 4

	fmt.Println("Your array is: ", myArr)
	fmt.Println("With length of: ", len(myArr))

}
