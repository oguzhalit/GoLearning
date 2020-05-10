package main

import "fmt"
// fonksiyon tanımlamak için func kullaniyoruz.void static public,private yok.Mantıgı var.
// go tool'ları ile sytanx ve standartı check eden durumları engeller.
// kullanmakta fayda var
func main() {
	fmt.Println("Merhaba, Dunya")
	// degisken tanimlama yöntemleri
	var merhaba string = "dunya"
	// tipini degisken isminden snra veriyoruz
	var mars = "dunya" // degiskene bakarak tipini anla ve ata
	jupiter := "dunya" // jupiyer'i tanımlıyor, atatıgıma göre tpini belirle
	// kullanılmayan durumlar go tarafından silinir. ve gereksi kod kullanımını engeller
	// degiskenler deger atanmadıgında kendiniz zero haline getirirler
	// go'da null yok, nil var
	fmt.Println(mars,jupiter,merhaba)
	var sayi int
	fmt.Println(sayi) //burda 0 cıktı vermesinin nedeni degeri olmaması, deger atamanıza gerek kalmadan zero degerini alıyor.

}
