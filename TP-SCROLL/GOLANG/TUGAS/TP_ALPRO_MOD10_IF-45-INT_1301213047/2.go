package main

import . "fmt"

func linearSortingVal(myArray []int, n int) int {
	for i, v := range myArray {
		if n == v {
			return i
		}
	}
	//
	return -1
}

func main() {
	n := []int{9, 1, 33, 21, 78, 42, 4}

	Println(linearSortingVal(n, 78)) // 4
}
