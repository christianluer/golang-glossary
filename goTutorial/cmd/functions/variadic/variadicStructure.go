package variadicStructure

// arguments are multiples only works if is the last argument with ...type, and to unfurl it is type...
func AddNumbers(ints ...int) int { // this means it will take any arguments 1,2,3,4,5,6... and it will convert it to []int
	n := 0
	for _, v := range ints {
		n += v
	}
	return n
}

func UnfurlingSlice(ii []int) int {
	return AddNumbers(ii...)
}
