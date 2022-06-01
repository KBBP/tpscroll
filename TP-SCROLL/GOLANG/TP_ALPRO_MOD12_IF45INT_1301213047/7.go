package main

func sortedB(tab [9999]int, n int) {
	var i, x, k int
	for i = 1; i < n; i++ {
		k = tab[i]
		x = i - 1

		for x >= 0 && k > tab[x] {
			tab[x+1] = tab[x]
			x -= 1
		}
		
		tab[x+1] = k
	}
}
