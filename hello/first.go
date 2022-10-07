package main
import "fmt"

// jwtToken := 300000      not allowed to use this operator outside function

var jwtToken = 300000

const LoginToken string = "bfvbhfbvhjdbh"  // public

func main() {
	fmt.Println("Hello from Learncodeonline.in")
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.41558541584551415848151
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVarible int
	fmt.Println(anotherVarible)
	fmt.Printf("Variable is of type: %T \n", anotherVarible)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	// website = 3

	// no var style 
	numberOfUser := 300000.23
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}