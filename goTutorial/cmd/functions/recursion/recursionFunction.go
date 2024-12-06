package recursion

import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func CallFactorial(n int) {
	fmt.Println(Factorial(n))
}
