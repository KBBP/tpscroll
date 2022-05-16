package main

import . "fmt"

type arr [256]int

func rev(a arr, b *arr) {
	var (
		begin, end, temp int
	)
	begin = 0
	end = len(a) - 1
	for begin < end {
		temp = a[begin]
		b[begin] = a[end]
		b[end] = temp
		begin = begin + 1
		end = end - 1
	}
}

func palindrome(a arr) bool {
	var (
		start, end int
		pal        bool
	)
	pal = true
	start = 0
	end = len(a) - 1
	for pal == true && start < end {
		if a[start] == a[end] {
			pal = true
		} else {
			pal = false
		}
		start = start + 1
		end = end - 1
	}
	return pal
}

func fill(a *arr) {
	var end, x int
	end = len(a)
	for i := 0; i < end; i++ {
		Scan(&x)
		a[i] = x
	}
}

func main() {
	var (
		T1, T2 arr
		pal    bool
	)
	T1 = [256]int{10, 20, 30, 30, 20, 10}

	rev(T1, &T2)
	Println(T2)

	pal = palindrome(T1)

	Println(pal)

	fill(&T1)
	fill(&T2)

	Println(T1)
	Println(T2)
}
