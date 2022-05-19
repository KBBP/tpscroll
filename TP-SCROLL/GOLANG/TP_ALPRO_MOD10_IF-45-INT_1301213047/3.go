package main

import . "fmt"

func binSearch(myArray []int, s int) int {
	var (
		mid, midVal int
	)

	lefty := 0
	righty := len(myArray) - 1 // GETTING LENGTH

	for lefty <= righty {
		mid = (lefty + righty) / 2
		midVal = myArray[mid]

		if midVal == s {
			return mid
		} else if midVal < s {
			lefty = mid + 1
		} else {
			righty = mid - 1
		}
	}

	return -1
}

func main() {
	n := []int{2, 9, 11, 21, 22, 32, 36, 48, 76}

	Println(binSearch(n, 32))
}
