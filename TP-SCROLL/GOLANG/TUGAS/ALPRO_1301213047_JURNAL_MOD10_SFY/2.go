package main

import . "fmt"

var tempArray [20]int
var secArray [20]int

func main() {
	var userInput, suaraMasuk, suaraSah, max, temp int

	suaraMasuk = 0
	suaraSah = 0

	Println("Please enter input: ")
	Scan(&userInput)

	for userInput != 0 {
		suaraMasuk += 1
		if userInput >= 1 && userInput <= 20 {
			suaraSah += 1
			tempArray[userInput-1] += 1
		}

		Scan(&userInput)
	}

	Println("Suara masuk: ", suaraMasuk)
	Println("Suara sah: ", suaraSah)

	for i := 0; i < 20; i++ {
		if tempArray[i] != 0 {
			Println(i+1, ":", tempArray[i], "suara")

			if tempArray[i+1] >= tempArray[i] {
				max = tempArray[i+1]
				secArray[i] = tempArray[i+1]
			}
		}
	}

	Println("Current highest standing: ", max, "suara")
	Println("secArray:", secArray)

	for k := 0; k < 20; k++ {
		if secArray[k] != 0 {
			if secArray[k+1] >= secArray[k] {
				temp = secArray[k+1]
			}
		}
	}

	Println("recap temp: ")
	Println(temp)
}
