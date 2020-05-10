package main

import "fmt"

type accountNumber int

func main() {
	fmt.Println(ismi("oguz","sak"))
}
// birden fazla type dondurebilirm, fonksiyonla
func ismi(firstName string, lastName string) string{
	return firstName + lastName
}
