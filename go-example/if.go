package main

import "fmt"

func main(){
	if 7 %2 == 0 { 
		fmt.Println("7 sayisi çift") 
	}else { // burda cıkıntılara dikkat etmek laazim
		fmt.Println("7 sayisi tek")
	}

	if 8%4 == 0 {
		fmt.Println("8 sayisi 4'e tam bölünür.")
	}

	if num :=9; num < 0{
		fmt.Println(num,"negatif")
	}else if num < 10{
		fmt.Println(num,"tek haneli")
	}else {
		fmt.Println(num, "cift haneli")
	}
}
