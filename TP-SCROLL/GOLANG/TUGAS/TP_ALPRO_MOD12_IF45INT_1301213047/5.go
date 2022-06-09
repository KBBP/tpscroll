package main

import . "fmt"

func main() {
	var (
		T    = []int{1, 2, -9, -7, -15, 9999}
		v, n int
	)
	v = 2
	n = len(T)
	addData((*[9999]int)(T), &n)
	n = len(T)
	Println(findSeq(T, v))
	Println(minVal(T, n))
}
