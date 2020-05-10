package main

//metot bir sınıf ile ilişkilendirilen bişey
//diger dillerdeki gibi sınıf yok, benzer bir yapı var

import "fmt"
// type ile tip tanımlayalım
type Account struct{
	AccountNumber int
	FirstName string
	LastName string
	//filled tanımyalalım.
}
// account ile tipini tanmladık
func (account Account) GetFullName() string{
//return olarak string döndürecek
//analoji ve metotloji
	//account.FirstName = "ali" //override edersek nolcak bakalım
	return account.FirstName + " " + account.LastName
}

func (account Account) ChangeFirstName(firstName string) {
	account.FirstName = firstName //bu şekilde oveeride deneyelim
}
// buna metot tanımlayalım
func main () {
	account := Account {2020,"Oguz","Sak"}
	account.ChangeFirstName("Hakan") //halen Oguz sak yazdırıyor. bunun nedeni deger kopyaları olusturuyor.yani ayni accountu gösterrmiyor.Pointer seklinde degistirirsek aynı olarak etkilenir.Aynı objeye ulaşma özelligi kazandrıyor.
	 
	fmt.Println(account.GetFullName())
	var accountNumber int = 500
	var number *int = &accountNumber
	fmt.Println(*number)//degerin bulundugu yere işaret ederek, kopyalama ve tekrar tanımlama durumlarından bizi kurtarıyor.
}
