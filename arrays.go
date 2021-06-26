package main

import (
	"fmt"
)

func main() {
	//grades := [3]int { 96, 91, 86}
	grades := [...]int {98, 81, 71}
	fmt.Printf("Grades : %v\n", grades)
	var students [3]string
	students[0]= "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Num. Students: %v\n", len(students))
	//var identityMat [3][3]int = [3][3]int{[3]int{0, 0, 1}, [3]int{0,1,0}, [3]int{1,0,0}}
	//var identityMat = [3][3]int{{0, 0, 1}, {0,1,0}, {1,0,0}}
	identityMat := [3][3]int {{0, 0, 1}, {0,1,0}, {1,0,0}}
	fmt.Println(identityMat)
//	arrays are considered as values. they are copied by value
	a:= [...]int {1, 4, 3}
	b := a
	b[1] = 2
	fmt.Println(a);
	fmt.Println(b);
//	to copy address instead
	c := &a
	c[0] = 9
	fmt.Println(a)
	fmt.Println(c)
//	slices
	s := []int {1,2,3}
	fmt.Printf("Length: %v\n", len(s))
	fmt.Printf("Capacity: %v\n", cap(s))
	t := s
	t[2] = 5
	fmt.Printf("s: %v\n", s)
	fmt.Printf("t: %v\n", t)
	ten := []int{1,2,3,4,5,6,7,8,9,10}
	all := ten[:]
	some := ten[3:8]
	tenArr := [...]int{1,2,3,4,5,6,7,8,9,10}
	all = tenArr[:]
	some = tenArr[3:8]
	fmt.Println(all)
	fmt.Println(some)
//	making slices
	slice := make([]int, 3, 10)
	fmt.Printf("Length: %v\n", len(slice))
	fmt.Printf("Capacity: %v\n", cap(slice))
	slice = append(slice, 3, 8, 9)
	slice[0] = 3
	slice[1] = 4
	slice[2] = 5
	fmt.Println(slice)
//	spread operator
	slice = append(slice, []int{5, 7}...)
//	removing elements
	slice = slice[2:]
	slice = slice[:len(slice) - 2]
	fmt.Println(slice)
	slice = append(slice[:2], slice[3:]...)
	// This append slices way to remove elements from a to get b messes with a. Unexpected behavior.
	//b := append(a[:2], a[3:]...)
	fmt.Println(slice)
}
