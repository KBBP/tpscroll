package main

func sortedA(tab *[9999]int, n int) {
	var i, x, k int
	for i = 1; i <= n; i++ {
		k = i
		for x = i + 1; x < n; x++ {
			if tab[x] < tab[k] {
				k = x
			}
		}
		if tab[i] != tab[k] {
			replace(&tab[i], &tab[k])
		}
	}
}

func replace(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
