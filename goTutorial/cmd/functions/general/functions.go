package general

import "fmt"

// Defer method
func FunctionToBeDefered() {
	fmt.Println("defered")
}
func ThisIsCalledSecond() {
	fmt.Println("This is called second but the first is defered to the end")
}

func DeferFunction() {
	defer FunctionToBeDefered() // this makes it execute at the end of the braces of the function thats being called
	ThisIsCalledSecond()
}

////////
