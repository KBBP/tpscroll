package main

import . "fmt"

func shaggy(tab [9999]trec, n int) {
	var (
		index     int
		f1, f2, f bool
	)

	index = 2
	f = false

	for index < n && !f {
		f1 = tab[index-1].v1 == tab[index].vx.v2
		f2 = tab[index].vx.v3 == tab[index].v4
		f = f1 && f2
		index += 2
	}

	if f {
		Println("There is Shaggy there. Which is…")
	}
}
