package main

import (
	"fmt"
)

func main() {
	//	TODO: MAKE A RECURSIVE ARRAY
	var arr = []int{1, 3, 5, 7, 9, 9, 7, 1, 3}

	displayArray(arr, 0)
}

func displayArray(arr []int, startIndex int) {
	if startIndex < len(arr) {
		fmt.Print(arr[startIndex], " ")
		displayArray(arr, startIndex+1)
	} else {
		fmt.Println(" ")
	}
}
