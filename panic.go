package main

import (
	"fmt"
	"log"
)

func main() {
// 	typical usage
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello Go!"))
	//})
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err.Error())
	//}

//	recover from panic example

	fmt.Println("start")
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Println("Caught error: ", err)
	//	}
	//}()
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		//	if you decide that the error is unrecoverable, simply panic again
		//	panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}