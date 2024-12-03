package object

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	kWh       uint8
	mpkWh     uint8
	ownerInfo owner
}

type owner struct {
	name string
	age  uint8
}

func (g gasEngine) milesLeft() int8 {
	return int8(g.mpg * g.gallons)
}

func (e electricEngine) milesLeft() int8 {
	return int8(e.mpkWh * e.kWh)
}

type engine interface {
	milesLeft() int8 // method signature for milesLeft is an abstract? method
}

func object() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex", 25}}
	fmt.Println(myEngine)
	var myStruct = struct {
		name string
		age  uint8
	}{"haha", 60}
	fmt.Println(myStruct)
	fmt.Printf("driving myEngine: %v\n", myEngine.milesLeft())
	fmt.Printf("Can myEngine make it 100 miles? %v\n", canMakeIt(myEngine, 100))
	var myElectricEngine electricEngine = electricEngine{4, 30, owner{"Alex", 25}}
	fmt.Printf("driving myElectricEngine: %v\n", myElectricEngine.milesLeft())
	fmt.Printf("Can myElectricEngine make it 100 miles? %v\n", canMakeIt(myElectricEngine, 100))
}

func canMakeIt(g engine, miles int8) bool {
	return g.milesLeft() > miles
}
