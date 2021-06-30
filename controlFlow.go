package main

import "fmt"

func main() {
	statePopulations := map[string]int {
		"California" : 39250017,
		"Texas" : 27862596,
	}
	// pop, ok live within the if-chain, it doesn't exist outside.
	// note that we can redeclare them using :=
	if pop, ok := statePopulations["Ohio"]; ok {
		fmt.Println(pop)
	} else if pop, ok = statePopulations["California"]; ok {
		fmt.Println(pop)
	}
//	logical operators: ||, &&, !

	switch i:=2+3; i {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 4, 7:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
	num := 10
	// in Golang, the break is implied. Not implicit fallthrough like C-based code.
	switch {
	case num <= 10:
		fmt.Println("num less than equal 10")
	case num <= 20:
		fmt.Println("num less than equal to 20")
	default:
		fmt.Println("num greater than 20")
	}
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is type int")
		if i.(int) > 0 {break}
		fmt.Println("unwanted output")
	case float64:
		fmt.Println("i is type float64")
	case string:
		fmt.Println("i is type string")
	case [3]int:
		fmt.Println("i is length 3 int array")
	default:
		fmt.Println("i is another type")
	}
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	j := 5

	for j < 10 {
		j++
	}

	for {
		if j > 56 {break}
		j *= 2
	}

	s := [...]int {1, 8, 4, 6}
	for k, v := range s {
		fmt.Println(k, v)
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
//	for loop over channels, used for concurrent programming
}