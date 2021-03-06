package main

import "fmt"

type Account struct{
	AccountNumber int
	FirstName string
	LastName string
}

func main(){
	account:=Account{100,"Okyanus","Kuce"}
	ChangeAccountNumber(account)
	fmt.Println(account.AccountNumber)
	fmt.Printf("Main function icerisinde ki deger %p \n",&account)
	ChangeAccountNumberWithPointer(&account)
	fmt.Println("Now account number after pointer : ",account.AccountNumber)
}

func ChangeAccountNumber(account Account){
	account.AccountNumber = 120
	fmt.Printf("Method icerisinde ki deger %p \n",&account)
}

func ChangeAccountNumberWithPointer(account *Account){
	account.AccountNumber = 200
	fmt.Printf("Pointer ile gelen method icerisinde ki deger %p \n",account)
}