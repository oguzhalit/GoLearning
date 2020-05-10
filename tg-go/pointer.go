package main

import "fmt"

type Account struct{
	AccountNumber int
	FirstName string
	LastName string
}

func main() {
	account := Account{132,"Oguz","Sak"}
	ChangeAccountNumber(account)
	fmt.Println(account.AccountNumber)//accountlar arası fark var
	fmt.Printf("main: %p \n",&account)
	//var acountNumber int = 300
	//pointer'e accountNumber'ın adresini alyrum
	//var accountNumberPointer *int = &acountNumber
	//fmt.Println(accountNumberPointer)
	//fmt.Println(*accountNumberPointer)
}
// pointer şeklinde degiştirme yapsaydık bu sefer override yapardık
func ChangeAccountNumber(account Account){
	account.AccountNumber = 234
	fmt.Printf("Change: %p \n ",&account)
}
