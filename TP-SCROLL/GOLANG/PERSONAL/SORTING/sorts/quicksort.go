package insertionsort

import (
	"fmt"
	"strconv"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		var pivot = partition(arr, low, high)
		quickSort(arr, low, pivot)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	var (
		pivot = arr[low]
		i     = low
		j     = high
	)

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot

	return j
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}

	fmt.Println("")
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

	quickSort(arr, 0, len(arr)-1)
	fmt.Print("After Sorting: ")
	printArray(arr)
}
