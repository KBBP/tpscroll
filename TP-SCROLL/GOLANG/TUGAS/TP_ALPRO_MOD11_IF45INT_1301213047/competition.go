package main

import . "fmt"

// Penggunaan titik untuk alias sebagai FMT agar bisa dihilangkan

func main() {
	var sizeArray int

	Printf("Enter size of your array: ")
	Scanln(&sizeArray)

	var myArr = make([]int, sizeArray)

	for i := 0; i < sizeArray; i++ {
		Printf("Enter %dth element: ", i)
		Scan(&myArr[i])
	}

	// Test Test
	Println("Your array before sort is: ", myArr)
	Println("Your array after sort is: ", sortFunc(myArr))
	Println("Total sum of array: ", sumMedian(myArr))
	Println("With length of: ", len(myArr))
}

func sortFunc(myArr []int) []int {
	for k := 0; k < len(myArr)-1; k++ {
		indexMin := k
		for x := k + 1; x < len(myArr); x++ {
			if myArr[indexMin] > myArr[x] {
				indexMin = x
			}
		}
		tempo := myArr[k]
		myArr[k] = myArr[indexMin]
		myArr[indexMin] = tempo
	}

	return myArr
}

func sumMedian(myArr []int) int {
	res := 0

	for _, v := range myArr {
		res += v
	}

	// Finding median
	if res%2 == 0 {
		//	EVEN
	} else {
		//	ODD
	}

	return res
}

// puzing maszeh :)
// ikhsan assidiqie property anyways
