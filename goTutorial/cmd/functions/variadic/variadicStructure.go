package variadicStructure

func AddNumbers(ints ...int) int {
	n := 0
	for _, v := range ints {
		n += v
	}
	return n
}

func UnfurlingSlice(ii []int) int {
	return AddNumbers(ii...)
}
