package main

import "fmt"

type Account struct{
	/*
		Class'lar yerine go'da bu sekilde type'lar tanimlanabiliyor
		Icerisinde method ya da baska herhangi bir yapi tanimlanamiyor
	*/
	AccountNumber int
	FirstName string
	LastName string
}

//GetFullName fonksiyonunu Account tipine tanimla
func (account *Account) GetFullName() string{
	return account.FirstName + " " + account.LastName
}

func (account *Account) ChangeFirstName(firstName string) {
	account.FirstName = firstName
}

func main(){
	//:= kullanimi variable'i hem create eder hemde initialize eder
	account:=Account{1,"Okyanus","Kuce"}
	account.ChangeFirstName("Cihangir")
	fmt.Println(account.GetFullName())
}
