package anonymous

import "fmt"

func RunAnonymousFunc() {
	func() {
		fmt.Printf("Im an anonymous func running by itself\n")
	}()

	func(s string) {
		fmt.Printf("Im an anonymous func showing my name: %s\n", s)
	}("Chris")

	functionVar := func(s string) {
		fmt.Printf("Im an anonymous func showing my name: %s\n", s)
	}
	functionVar("Chris")
}
