package sliceStructure

import (
	"errors"
	"fmt"
)

func SliceStructure(length int) []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i + 1
	}
	return slice
}

var SliceOne = []int{1, 2, 3, 4, 5}
var SliceTwo = []int{6, 7, 8, 9, 10}

func AppendToSlice(sliceArray []int, number int) []int {
	return append(sliceArray, number)
}

func AppendMultipleToSlice(sliceArray []int, numbers []int) []int {
	sliceArray = append(sliceArray, numbers...)
	return sliceArray
}

func PrintUsingRange(sliceArray []int) {
	for _, v := range sliceArray {
		fmt.Printf("value %v \n", v)
	}
	for i, v := range sliceArray {
		fmt.Printf("value indexed Index: %v, value: %v\n", i, v)
	}
}

func SliceTheSlice(sliceArray []int, startSlice, endSlice int) ([]int, error) {
	if startSlice > endSlice || startSlice < 0 || endSlice < 0 || startSlice > len(sliceArray) || endSlice > len(sliceArray) {
		return nil, errors.New("please enter start > end indexes that also are in the slice")
	}
	return sliceArray[startSlice:endSlice], nil
}

func DeleteFromTheSlice(sliceArray []int, index int) ([]int, error) {
	if index < 0 || index > len(sliceArray) {
		return nil, errors.New("please enter a valid index")
	}
	return append(sliceArray[:index], sliceArray[index+1:]...), nil
}

func MakeSliceWithCapacity(capacity int) {
	slice := make([]int, 0, capacity)
	fmt.Printf("Length: %v, Capacity: %v\n", len(slice), cap(slice))
}

func MultiDimentionalSlice(capacity int) ([][]int, error) {
	example := [][]int{{1, 2}, {3, 4}}
	fmt.Println(example)
	return make([][]int, capacity), nil
}

func CopyArray(copyItem []string) []string {
	copiedArray := make([]string, len(copyItem))
	copy(copiedArray, copyItem)
	return copiedArray
}
