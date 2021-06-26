package main

import "fmt"

const (	a = iota
	b
	c
)

const (	a2 = iota
)

const (
	errorType = iota
	catType
	dogType
	chickenType
)
const (
	_ = iota // throw away 0 please.
	carType
	trainType
	planeType
)

const (_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAmerica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a2)
//	what's the point of iota? use as a Enumeration
	var animalType int = dogType
	fmt.Println(animalType == catType)
//	better one
	var filesize float32 = 1230000
	fmt.Printf("%vMB\n", filesize/MB)
	var role byte = canSeeSouthAmerica | canSeeAsia | isHeadquarters
	fmt.Printf("%b\n", role)
}