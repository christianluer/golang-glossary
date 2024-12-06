package structStructure

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

type Adult struct {
	Person
	Job string
}

type Child struct {
	Person
	School string
}

func AdultStructure(name string, age int, job string) Adult {
	structure := Adult{
		Person: Person{
			Name: name,
			Age:  age,
		},
		Job: job,
	}
	return structure
}

func ChildStructure(name string, age int, school string) Child {
	structure := Child{
		Person: Person{
			Name: name,
			Age:  age,
		},
		School: school,
	}
	return structure
}

func PersonStructure(name string, age int, ocupation string) interface{} {
	license := []string{"Driver", "Pilot", "Doctor", "Developer"}
	if ocupation2 := ocupation; ocupation2 == license[3] {
		fmt.Println("Adult structure created and a developer")
		return AdultStructure(name, age, ocupation)
		// make ocupation 2 check if it is in the 0 to 2 range
	} else if age2 := age; age2 >= 18 {
		fmt.Println("Adult structure created")
		return AdultStructure(name, age, ocupation)
	} else {
		fmt.Println("Child structure created")
		return ChildStructure(name, age, ocupation)
	}
}

func PrintAnonimousStruct() {
	anonimousStruct := struct {
		name string
		age  int
	}{
		name: "hello test",
		age:  20,
	}
	fmt.Printf("type %T, %v", anonimousStruct, anonimousStruct)
}
