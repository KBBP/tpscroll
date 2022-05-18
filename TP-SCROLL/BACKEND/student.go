package main

// TODO: import base.json into student.go

import . "fmt"

type BaseArray struct {
	name string
	eprt int
	gpa  float64
	sem  int
}

// --> {name, gpa, eprt, sem} constructor

func main() {

	users := BaseArray{
		name: "John Doe",
		eprt: 550,
		gpa:  3.6,
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
		eprt: 550, // (array[index].eprt)
		gpa:  3.4, // (array[index].gpa)
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

func lowestGpa() int {
	users := BaseArray{
		name: "Maria Doe",
		eprt: 550,
		gpa:  3.6,
		sem:  2,
	}

	Println(users)

	return 0
}

func avgSem() {

}
