package callback

import "fmt"

func add(a int, b int) int {
	return a + b
}
func subsctract(a int, b int) int {
	return a - b
}

func doOp(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func Execute() {
	opSum := doOp(3, 4, add)
	opSubs := doOp(4, 3, subsctract)
	fmt.Println(opSum)
	fmt.Println(opSubs)
}
