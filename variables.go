package main

import (
	"fmt"
	"strconv"
)

// package level declare, must use full syntax
var i float32 = 32.34895
// user var grouping
var (
	//actorName string = "Lisa"
	companion string = "Joseph"
	doctorNumber int = 3
	season int = 11
)
/*
	Three levels of scope: global, package, and block level.
	Global is just capitalized naming, otherwise package.
*/
func main() {
	//k := "unused"
	/* Shadowing: declaring same variable again in the inner scope overshadows the outer scope
	*/
	j := 58
	fmt.Printf("%v, %T\n", doctorNumber, doctorNumber)
	strJ := strconv.Itoa(j)
	fmt.Printf("%v, %T\n", strJ, strJ)
}