package main

import "fmt"

type AccountNumber int

func main() {
	var accountNumber AccountNumber = 5
	var number = 5 // ikisi icinde aynı degerleri atadagııma ragmen
				   // tipler farklı, accountnumber ve telefonnumaarası aynı olmasına ragmen farklı tipler kullanarak bunların çakışmasını önleriz
	if(accountNumber == number) {
		fmt.Println("Esittir")
	} else {
		fmt.Println("Degildir")
	}
}
