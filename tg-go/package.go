package main

// go dilinde packagelere kütüphane denebilir.
// go get ile paket indirip kullanabiliriz.
// kullanılmayan paketleri import edilmesine izin vermez.
// import yeniad "fmt" diyerek adını degiştirebiliriz.
// artık fmt yerine yeniad yazarak paketi kullanabiliriz.
// bos import etme özelligi ile kullanlmayan paketleri koyabilirzsiniz.
// go list ... ile yüklü paketleri görebilirim
// go list ./... mevcut kodumuzdaki paketleri görebiliriz.
import (
	"fmt"
	"strconv"
	//"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Merhaba, Dunya" + strconv.Itoa(degisken))
	
}
