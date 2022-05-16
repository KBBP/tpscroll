package main

import . "fmt"

type Customer struct {
	name, address, occupancy, salary string
	age, height                      int
	married                          bool
}

func main() {
	ikhsan := Customer{
		name:      "Ikhsan Assidiqie",
		address:   "Bandung",
		occupancy: "Student",
		salary:    "3.000.000",
		age:       20,
		height:    175,
		married:   false,
	}

	Println(ikhsan)
}
