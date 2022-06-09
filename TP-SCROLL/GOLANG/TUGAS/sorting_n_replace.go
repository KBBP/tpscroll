package main

import "fmt"

func main() {
	var grade = []int{3, 1, 6, 30, 8, 9, 2, 5, 9} // Panjang Array Real 9
	var temp int

	for i := 1; i < len(grade); i++ {
		for k := 0; k < len(grade)-1; k++ {
			temp = grade[k]
			grade[k] = grade[k+1]
			grade[k+1] = temp
		}
		fmt.Println(grade)
	}

	fmt.Println("Hasil akhir: ")
	fmt.Println(grade)
}
