package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	var fname string
	fmt.Print("Enter your firstname: ")
	fmt.Scan(&fname)
	fmt.Print("Enter your lastname: ")
	var lname string
	fmt.Scan(&lname)
	fmt.Print("Your Full Name is " + fname + " " + lname)
}
