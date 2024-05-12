package main

import "fmt"

const LoginToken string = "gsjgdjdjkjkd" //public variable

func main() {
	var username string = "Suyash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.5764678757899
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
