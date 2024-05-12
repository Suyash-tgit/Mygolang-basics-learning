package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang, no super or parent

	suyash := User{"Suyash", "suyash@go.dev", true, 16}
	fmt.Println(suyash)
	fmt.Printf("suyash details are: %+v\n", suyash)
	fmt.Printf("Name is %v and email is %v.", suyash.Name, suyash.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
