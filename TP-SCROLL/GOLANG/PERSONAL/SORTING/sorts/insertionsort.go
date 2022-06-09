package insertionsort

import (
	"fmt"
)

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var (
			temp         = arr[i]
			holePosition = i
		)

		for holePosition > 0 && arr[holePosition-1] > temp {
			arr[holePosition] = arr[holePosition-1]
			holePosition--
		}

		arr[holePosition] = temp
	}
}

func main() {
	var size int

	fmt.Print("Enter you array size: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("Your array: ", arr)

	fmt.Print("Before Sorting: ")
	printArray(arr)

	insertionSort(arr)
	fmt.Print("After Sorting: ")
	printArray(arr)
}
