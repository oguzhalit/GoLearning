package main

import (
	"fmt"
	"time"
)

func main() {
	
	i := 2
	fmt.Print("Sayi ",i,"dır")
	switch i {
	case 1:
		fmt.Println("bir")
	
	case 2:
		fmt.Println("iki")
	case 3:
		fmt.Println("uc")
	}


	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("bugun haftasoonu")
	default:
		fmt.Println("bugun haftaici")
	}

	t := time.Now()
	switch {
	case t.Hour()<12:
		fmt.Println("ögleden once")
	default:
		fmt.Println("ögleden sonra")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("type %T", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	//switch yapısana tekrar dön
}
