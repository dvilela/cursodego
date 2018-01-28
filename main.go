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

	// Arrays and Slices... Aditional content: https://github.com/golang/go/wiki/SliceTricks
	var arr1 [5]int // an array. Arrays can NOT be changed. They are already initialized
	fmt.Println("arr1:", arr1, len(arr1), cap(arr1))

	var slice1 []int // a slice. Slices can be changed
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))

	slice2 := make([]string, 5) // declares and initialize the slice
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	cars := []string{"GTI", "Civic", "Focus", "Corolla"}
	fmt.Println("cars", cars, len(cars), cap(cars))

	// append to slice
	cars = append(cars, "Audi_S7")
	fmt.Println("S7 appended to cars:", cars, len(cars), cap(cars))

	// delete Corolla from slice
	cars = append(cars[:3], cars[4:]...)
	fmt.Println("Corolla cutted from cars:", cars, len(cars), cap(cars))

	// Interfaces
	jojo := model.Bird{"Jojo"}
	fmt.Printf("%s is saying something: %s\n", jojo.Name, doCluck(jojo))
	fmt.Printf("%s is saying something: %s\n", jojo.Name, doQuack(jojo))

}

func doSomething() (success bool, err string) {
	err = "This is not the droid you are looking for"
	return
}

func doCluck(c model.Chicken) string {
	return c.Cluck()
}

func doQuack(d model.Duck) string {
	return d.Quack()
}
