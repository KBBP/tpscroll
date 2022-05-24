package main

import "fmt"

//func sort(tabReceiver []int) {
//	const temp = 0
//
//	for i := 0; i < len(tabReceiver); i++ {
//		fmt.Println(tabReceiver[i])
//	}
//}

func main() {
	var inputNumber int
	var storeArray []int
	fmt.Scanln(&inputNumber)
	storeArray[0] = inputNumber

	for i := 1; i < len(storeArray); i++ {
		storeArray[i] = inputNumber
		fmt.Scanln(&inputNumber)
	}

	//sort(storeArray)
}
