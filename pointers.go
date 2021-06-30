package main

import "fmt"

type myStruct struct {
	foo int
}
func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	fmt.Println(&a, b)

	arr := [3]int {3, 4, 6}
	i1 := &arr[2]
	i2 := &arr[1]
	fmt.Printf("%p, %p\n", i1,i2)

	var sp *myStruct
	// GOOD HABIT: CHECK FOR NIL POINTERS
	fmt.Println(sp)
	sp = &myStruct{foo: 9}
	fmt.Println(sp)
	fmt.Println(*sp)
//	use new syntax, but can't initialize
	sp = new(myStruct)
	fmt.Println(sp)
	(*sp).foo = 77
//	syntactic sugar expression, equal to sp->foo in C/C++
	sp.foo = 88
//	dereferencing pointers (*) has lower precedence over the .
	fmt.Println((*sp).foo)

//	copying Slices and Maps is essentially copying pointers, not values like arrays or structs.
}