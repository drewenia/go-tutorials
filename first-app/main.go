package main

import "fmt"

type Account struct{
	AccountNumber int
	FirstName string
	LastName string
}

func main(){
	var accountNumber int = 300
	var accountNumberPointer *int = &accountNumber //& ampersand isareti degeri degil adresi verir

	fmt.Println(accountNumberPointer)
	fmt.Println(*accountNumberPointer)
}