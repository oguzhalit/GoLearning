package main

import(
	"fmt"
	"time"
)

func main(){
	paket := make(chan string)

	go Mesaj(paket, "enektarlar koltugun alnda kalık beni ara")
	go Alıcı(paket)
	var erkenInput string
	fmt.Scanln(&erkenInput)
	fmt.Println("anlasilmiyor")
}

func Mesaj(mesajbirak chan string, icerik string){	
	fmt.Println("Baglanılıyor....")
	time.Sleep(time.Second*3)	
	fmt.Println("Baglantı kuruldu.")
	mesajbirak<-icerik
}

func Alıcı(mesajbirak chan string){
	fmt.Print("Baglantı acildi\n")
	gelen:=<-mesajbirak
	fmt.Println(gelen)
}
