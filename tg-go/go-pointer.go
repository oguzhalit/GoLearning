package main

import "fmt"
//pointerlar go'da daha sade
//pointer işaret eden.
//birini eve çagırrısanız ona adresini verirsiiz.
//baskalarını çagırmak içinde aynı adresi veriririz.
//evimi sana getiriyim veya her birinize bir ev veriyrum demeyiz
//kendi adresimizize çagırırız
//memroy bir sehir olarak düşünürsek
//bi yere veri yollamak istedigimizde adres lazım olur(posta kodu gibi)
// tekrarlardan kaçınarak perfrmans arttırmak önemlidir.
// adresini paylaşarak tekrarlardan kaçınabiliriz.
// verinin aynısını kopyalıp göndermek yerine.
// verinin adresini tutan tip => pointer
func main() {
	sayi := 2 //sayi adresinde saklanan veri 2
	// adresi yollayarak bu veriyee ulaşmasını saglayabilirz.
	//deneme := "merhaba , dunya"
	var test2 *int = &sayi
	test := &sayi // int veri tipinden bir veri tanımla(isaret edilen verinin tipi int)
	                  // ve bana memeory adresini döndür.
	fmt.Println(test) // bir adres, bu adresde sayi saklaniyor.
	                  // gerçek veri işaretlenen adresde
	fmt.Println(*test)//
	fmt.Println(*test2)
	fmt.Println(test2)
	degiskenicerigidegistir(test2)
	fmt.Println(test2)
	fmt.Println(*test2)
}

func degiskenicerigidegistir(geliy *int){
	*geliy = 3
}
