package main

func minVal(tab []int, n int) int {
	var minimum int
	for i := 1; i <= n; i++ {
		if tab[i] < tab[minimum] {
			minimum = i
		}
	}
	return minimum
}
