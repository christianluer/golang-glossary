package interfaceStructure

import (
	"fmt"
	"log"
)

// adding interface to receiver

type receiver struct {
	name string
	age  int
}

type ReceiverInterface interface {
	ReceiverName() (string, error)
	ChangeName(name string) error
	ChangeAge(age int) error
	Print()
}

func NewReceiver(name string, age int) ReceiverInterface {
	return &receiver{name: name, age: age}
}

func (re receiver) Print() {
	fmt.Println(re)
}

func (re *receiver) ChangeAge(age int) error {
	re.age = age
	return nil
}

func (re *receiver) ChangeName(name string) error {
	re.name = name
	return nil
}

func (re *receiver) ReceiverName() (string, error) {
	return re.name, nil
}

// adding method and a Interface

type Person struct {
	Name string
}

func (p Person) Speak() { // method
	fmt.Println("hi Im a person calling a method my name is: ", p.Name)
}

func (sa SecretAgent) Speak() {
	fmt.Println("Hi Im a secret agent calling the method I have license to kill:", sa.Ltk)
}

type SecretAgent struct {
	Person
	Ltk bool
}

type Human interface {
	Speak()
}

type SpeakPerson string

func (spkPerson SpeakPerson) Speak() {
	log.Println("Hi Im just and string that can speak", spkPerson)
}

func SpeakInterface(hu Human) {
	fmt.Println("calling the interface to actionate the speak method")
	hu.Speak()
}
