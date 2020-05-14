package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	fmt.Println("Series 1") //ilk işlemi süresinde
	time.Sleep(2 * time.Second) // ilk işlem bitmesini bekliycek 
	fmt.Println("Series 2") // ikinci işlemi yapabilmek için

	go func() {
		fmt.Println("Series 1")
		time.Sleep(2 * time.Second)
		done <- true
	}()
	fmt.Println("Series 2")
	result := <-done
	fmt.Println(result)
	//analoji yazımlım kodlarında
	
}
