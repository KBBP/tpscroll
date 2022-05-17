package main

import . "fmt"

type MyArrays struct {
	name      string
	eprt, gpa float64
	sem       int
}

// --> {name, gpa, eprt, sem} constructor

func main() {
	ikhsan := MyArrays{
		name: "Ikhsan Assidiqie",
		eprt: 5.4,
		gpa:  2.8,
		sem:  4,
	}

	Println("Name:", ikhsan.name,
		"\nGPA:", ikhsan.gpa,
		"\nEPRT Score:", ikhsan.eprt,
		"\nCurrent Semester:", ikhsan.sem)
}

func highestEprt(customArray MyArrays) {
	
}

func lowestGpa() {

}

func avgSem() {

}
