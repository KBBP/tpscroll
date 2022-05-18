package main

import . "fmt"

// TODO: import base.json into student.go

type BaseArray struct {
	name      string
	eprt, gpa float64
	sem       int
}

// --> {name, gpa, eprt, sem} constructor

func main() {

	users := BaseArray{
		name: "John Doe",
		eprt: 5.4,
		gpa:  2.8,
		sem:  4,
	} // PROJECTING ARRAYS INTO BaseArray
	// PULLING FILES LOGIC

	Println("Name:", users.name,
		"\nGPA:", users.gpa,
		"\nEPRT Score:", users.eprt,
		"\nCurrent Semester:", users.sem)
}

func highestEprt(N int) int {
	i := 1

	users := BaseArray{
		name: "John Doe",
		eprt: 5.4, // (array[index].eprt)
		gpa:  2.8, // (array[index].gpa)
		sem:  4,   // (array[index].sem)
	} // PROJECTING ARRAYS INTO BaseArray
	// PULLING FILES LOGIC

	max := users.eprt
	for i < N-1 {
		if users.eprt > max {
			max = users.eprt
		}
		i++
	}

	return int(max)
}

func lowestGpa() int{
	i := 1
	users := BaseArray {
		name: "Maria Doe",
		eprt: 6.6,
		gpa:2.6,
		sem: 2,
	}

	min := users.gpa
	for i < N-1 {
		if users.gpa < min {
			min = users.gpa
		}
		i++
	}

	return int(min)
}

func avgSem() {

}
