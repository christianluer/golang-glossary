package mapStructure

import (
	"errors"
	"fmt"
)

func CreateMap(name string, age int) (map[string]int, error) {
	newMap := make(map[string]int)
	newMap[name] = age
	newMap["test"] = 2
	return newMap, nil
}

func PrintMap(toPrint map[string]int) {
	fmt.Println(toPrint)
	// fmt.Printf("map to print: %#v\n", toPrint)
	// for k, v := range toPrint {
	// 	fmt.Printf("key %s, value: %d\n", k, v)
	// }
}

func DeleteMapElement(toDelete map[string]int, name string) error {
	// if toDelete[name] == 0 {
	// 	return errors.New("unsuported")
	// }
	if _, ok := toDelete[name]; !ok {
		fmt.Println("key didnt exist")
		return errors.New("unsuported")
	}
	delete(toDelete, name)
	return nil
}
