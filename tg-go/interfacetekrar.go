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
// tipinin ne oldugunu bilmiyorsak, type testini bos interfacces kullanarak gerçekleştirebilirz
// 

type CheckingAccount struct{
	AccountNumber int
}

// duck typing eger bir hayvan ordek gibi yürüyor ve ordek gibi ses cıkarıyorsa
// benim için o ordektir.

// bir account objesi Iaccount gibi getfullname döndrüyorsa benim 
// gözümde bu ıaccountdur.

func (account Account) GetFullName() string{
	return account.FirstName + " " + account.LastName //account üzerinden tanmladım
	// bu artık Iaccount gibi calisacaktir.
}

func (account Account) GetAccount() int{
	return account.AccountNumber
}

func printFullName(account IAccount){
	fmt.Println(account.GetFullName())
}
func main(){
	// Iaccount tipindede interfaces tanımı yapılabilir.implement edilebilir.
	//var myAccount IAccount = Account {}
	account := Account{"Oguz","Sak",12}
	printFullName(account)
	var name = "oguzsak"
	printScreen(name)

}

func printScreen(value interface{}){ //bos bir interfacces,her tipi uymmus oluyor bu sekilde
	hello, ok := value.(string)		// value ne oldugunu bilmiyoruz, ilerde string oldugunu ögrendik 
	if ok{
		fmt.Println(hello) //cevirilmişse sıkıntı yok
	}else{                 //value.(int) olarak donusturseydik hata alırdık.
		fmt.Println("olmadi") //interface tuple olarak düsünülebilir, degerin kendisi ve degerin tipini saklıyor.interfaces ile tipleri ögrenbilirsiiz.
	}
}									//diyelim. value.(string) dersek bulaiblirz.
