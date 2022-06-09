package main

import . "fmt"

func addData(tab *[9999]int, n *int) {
	var (
		data int
	)

	Scanln(&data)
	for data != 9999 {
		tab[*n] = data
		*n++
	}
}
