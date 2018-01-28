package main

import (
	"fmt"

	"github.com/dvilela/cursodego/government"

	"github.com/dvilela/cursodego/mathdv"
	"github.com/dvilela/cursodego/model"
)

// Go auto initializes the variables. They're not null!
// The package variables starting with Uppercase letter are public,
// and those with Lowercase are private.
var (
	// Name is the name of the person
	Name string // name = ""
	// Age is the age of the person
	Age int8 // age = ?
	// Married represents its civil state
	Married bool // married = false
	// private var
	id string // id = ""
)

func main() {
	// Variables
	msg := "Hello, World!"
	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("name: %s\n", Name)
	fmt.Printf("age: %d\n", Age)
	fmt.Printf("married: %t\n", Married)
	fmt.Printf("id: %v\n", id) // v can support multiple types
	// Functions and Loops
	for i := 0; i <= 10; i++ {
		fmt.Printf("Factorial of %d: %d\n", i, mathdv.Factorial(uint64(i)))
	}
	msg = "Hello, Range!"
	for i, letter := range msg {
		fmt.Printf("%q", letter)
		if i == len(msg)-1 {
			fmt.Println()
		}
	}
	// Structs
	john := model.Person{
		Name:    "John Doe",
		Married: true,
		Age:     31,
	}
	fmt.Printf("[Created] Person 1: %+v\n", john) // +v prints the vars names too
	// Pointers
	government.RegisterPerson(&john)
	fmt.Printf("[Registered] Person 1: %+v\n", john)
	// Conditionals
	if success, err := doSomething(); !success {
		fmt.Println("[Error] Could not do something. Cause:", err)
	}
}

func doSomething() (success bool, err string) {
	err = "This is not the droid you are looking for"
	return
}
