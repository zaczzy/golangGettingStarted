package main

import "fmt"

func main() {
	var b bool
	fmt.Printf("%v, %T\n", b, b)
	var (
		i1 int
		i2 int8
		i3 int16
		i4 int32
		i5 int64
	)
	var u1 uint16
	var by byte
	fmt.Printf("%v, %T\t" +
		"%v, %T\t" +
		"%v, %T\t" +
		"%v, %T\t" +
		"%v, %T\n", i1, i1, i2, i2, i3, i3, i4, i4, i5, i5)
	fmt.Printf("%v, %T\t" + "%v, %T\n", u1, u1, by, by)

//	bit wise operations
	a := 10
	c := 3
	fmt.Println(a & c) // and
	fmt.Println(a | c) // or
	fmt.Println(a ^ c) // xor
	fmt.Println(a &^ c) //andnot (set if neither input is set, the opposite of or)
//	bit shift
	fmt.Println(a << 3)
	fmt.Println(a >> 3)
	s := "this is a string"
	bs := []byte (s)
	bs[3] = 'c'
	fmt.Println(bs)
	fmt.Println(string(bs))
}