// interfaces
package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// the beauty here is that we can implement interface to match up concrete types, so those types automatically implements your interface. We don't have to design the interface at design time.


func main() {
	// polymorphic behavior!
	var w Writer = ConsoleWriter{}
	
	w.Write([]byte("hello go!"))
}