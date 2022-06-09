package main

func quickSearch(tab [9999]int, n, v int) int {
	l := 0
	r := n - 1
	mid := 0

	for tab[l] > tab[r] {
		mid = (l + r) / 2
		if v < tab[mid] {
			l = mid + 1
		} else if v > tab[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
