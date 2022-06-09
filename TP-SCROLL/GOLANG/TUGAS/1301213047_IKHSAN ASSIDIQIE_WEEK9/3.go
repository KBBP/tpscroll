package main

import . "fmt"

type TabGoal [100]int

func main() {
	var (
		T1, T2, T3 TabGoal
		n1, n2, n3 int
	)
	FillTab(&T1, &n1)
	FillTab(&T2, &n2)
	FillTab(&T3, &n3)
	Println(avg(T1, n1), avg(T2, n2), avg(T3, n3))
}

func FillTab(t *TabGoal, n *int) {
	var win int
	Scan(&win)
	*n = 0
	for win >= 0 {
		t[*n] = win
		*n++
		Scan(&win)
	}
}

func avg(t TabGoal, n int) float64 {
	var (
		x       int
		average float64
	)
	for i := 1; i <= n; i++ {
		x += t[i]
	}

	average = float64(x / n)
	return average
}
