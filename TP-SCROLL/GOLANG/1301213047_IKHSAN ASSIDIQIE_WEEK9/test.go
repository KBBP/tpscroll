package main

import . "fmt"

type Customer struct {
	name, address, occupancy, salary string
	age, height                      int
	married                          bool
}

func dataCustomer(sean string) int {
	return 0
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
	dio := Customer{
		name:      "Dio Irsaputra Siregar",
		address:   "Bandung",
		occupancy: "Student",
		salary:    "2.000.000",
		age:       19,
		height:    172,
		married:   false,
	}

	Println(ikhsan)
	Println(dio.name)
}
