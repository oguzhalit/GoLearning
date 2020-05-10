package main

import "fmt"

type IAccount interface{
	GetFullName() string
	GetAccount() int
}

type Account struct{
	FirstName string
	LastName string
	AccountNumber int
}
 
// duck typing

func (account Account) GetFullName() string{
	return account.FirstName + " " + account.LastName
}

func (account Account) GetAccountNumber() int{
	return account.AccountNumber
}

func (account Account) GetAccount() int{
	return account.AccountNumber
}

func printFullName(account IAccount){ //kücük harf private
	fmt.Println(account.GetFullName())
}

func main() {
	var myAccount IAccount = Account {"oguz","halit",123}
	account := Account{"Oguz","Sak",123}
	printFullName(account)
	fmt.Println(myAccount.GetFullName())
	var name = "oguz"
	printScreen(name)
}

func printScreen(value interface{}){
	hello, ok := value.(string)//tipi birebir dönüştürme yaptk
    if ok{																		
		fmt.Println(hello)
	}else{
		fmt.Println("OLmadi")
	}
}
