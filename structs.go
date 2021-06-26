package main

import (
	"fmt"
	"reflect"
)
type Doctor struct {
	number int
	actorName string
	companions []string
}
type Animal struct {
	Name string `required max:"100"`
	Origin string
}

// Bird tries to emulate an "is-a" relationship
// but still an independent struct type
// it's a "has-a" relationship
// When using Polymorphism, wanting the type to behave like the base-class
// Use Interface instead.
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}
func main() {
	statePopulations := map[string]int {
		"California" : 39250017,
		"Texas" : 27862596,
	}
	fmt.Println(statePopulations)
//	array is a valid key type, a slice is not
	m := make(map[[4]int]string)
	m = map[[4]int]string{
		{1, 4, 5, 6} : "value",
	}
	fmt.Println(m)
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Texas")
	population, ok := statePopulations["Californa"]
	fmt.Println(population, ok)
	population, ok = statePopulations["Georgia"]
	fmt.Println(population, ok)

	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])
//	anonymous struct
	anotherDoctor := struct{name string} {name: "John Pertwee"}
	fmt.Println(anotherDoctor)
//	structs are also copied by value, can use & for copies by address
	docMap := map[struct{name string; age int}]string{
		{name: "John", age: 38} : "Second floor",
		{name: "James", age: 29} : "First floor",
	}
	fmt.Println(docMap)
	b := Bird{}
	b.Name = ""; b.Origin = ""
//	tags, reflection
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // tags is just a string in themselves
}
