package main

import . "fmt"

type CustomerData struct {
	fullName, fullAddress, posCode, phoneNumber string
	distributingOption                          int
}

// Struct just like a Class in Java

func dynamicCallOut(name, address, pos, phone string, distributing int) CustomerData {
	dynamicSortOf := CustomerData{
		fullName:           name,
		fullAddress:        address,
		posCode:            pos,
		phoneNumber:        phone,
		distributingOption: distributing,
	}

	return dynamicSortOf
}

func main() {
	var (
		name, address, pos, phone string
		distributing, pickMenu, n int
	)

	Println("Welcome to 141's Armoury.")
	Println("What do you want to be served today?")
	Println("1. Submit Form \n2. Exit")

	Scanln(&pickMenu)
	Scanln(&n)

	for i := 1; i <= n; i++ {
		// LOOPING HOW MANY CLIENTS
	}

	if pickMenu == 1 {
		Println("Please fill this form belo: ")
		Println("Format:\nName, Address, Postal Code, Phone, Distributing Option")
		Scanln(&name, &address, &pos, &phone, &distributing)
	} else if pickMenu == 2 {

	} else {
		Println("Invalid input.")
	}

	Println(dynamicCallOut(name, address, pos, phone, distributing))
}
