package main

import (
	"fmt"
	"sync"
)
var waitGroup = sync.WaitGroup{}
//var m = sync.RWMutex{}
// m.RLock()
// m.RUnlock()
// m.Lock()
// m.Unlock()

func main() {
	waitGroup.Add(1)
	// spin off green thread
	msg:="hello"
	go func(msg string) {
		fmt.Println(msg)
		waitGroup.Done()
	}(msg)
	msg = "goodbye"
	waitGroup.Wait()

}