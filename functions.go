package main

import "fmt"

func sayMessage(msg string) {
	fmt.Println(msg)
}
// syntax sugar params
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "ted"
}
func sum(msg string, values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
	return result
}
func localReturn () *int {
	localVar := 55
	return &localVar
}
//	weird syntax sugar : named return variable
func sumAgain(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}
// Multiple returns, useful for idiomatic use of error return
func divide(a, b float64) (float64, error) {
	if (b == 0.0) {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	sayMessage("Hello go!")
//	strings are passed by value, by passing pointers, we can modify the input
	name := "stacy"
	greeting := "hello"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
//	pass pointers is faster than passing large structs or arrays by value
//  passing slices or maps will always be by reference since they are represented by internal pointers

// variadic parameters
	sum("the sum is", 1,2,3,4,5)
//	GO MAGIC: allowing returning pointers to local variables
// HOW IT WORKS: automatic promotion of the return value to move from the execution stack to the heap,
// when runtime sees that you are returning a pointer on the stack.
	ref := localReturn()
	fmt.Println("the value is", *ref)
	fmt.Println(sumAgain(1,2,3,4))

//	error checking
	d, err:= divide(6, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
//	anonymous function
//	(immediately invoked)
	func(){fmt.Println("hello go")}()
// functions as types

	var div func(float64, float64) (float64, error)
	div = func(a, b float64) (float64, error) {
		if( b == 0.0) {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}
	var r float64
	r, err = div(4, 5)
	fmt.Println(r)
//	methods are functions with a type as a context
//	value receiver or pointer receiver
}
type greeter struct {
	greeting string
	name string
}
// value receiver
// cannot modify context
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
// pointer receiver
func (g *greeter) modGreet(greeting string) {
	g.greeting = greeting
}