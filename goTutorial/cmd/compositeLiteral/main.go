package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string // if you add capital letter it will be exported outside the package
}

func main() {
	// composite literal is a way to declare a value for a composite type: type{values}
	x := []int{1, 2, 3, 4, 5} // here x is a slice of int so it is a composite type []int and {}
	fmt.Printf("%T\n", x)

	m := map[string]int{"Todd": 45,
		"Job": 42,
	} // here m is a map of string to int so it is a composite type map[string]int and {}
	fmt.Printf("%T\n", m)
	fmt.Println(m)
	p1 := person{
		"James",
		"Bond",
	}
	fmt.Println(p1)
}
