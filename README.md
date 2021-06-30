## Notes
## Credit: freeCodeCamp.org tutorial
https://www.youtube.com/watch?v=YS4e4q9oBaU
### Variable declaration
- three ways to declare
    - var foo int
    - var foo int = 42
    - foo := 42
- can't redeclare variables, but we can't shadow them
- all variables must be used
- Visibility
    - lower case first letter for package scope
    - upper case first letter to export
    - when in a block, block scope, no private scope.
- Naming convention
    - Pascal or camelCase
        - Captialize acronyms (HTTP, URL)
    - as short as possible
        - longer names for longer lives
- Type conversion
    - func like, destinationType(variable)
    - no implicit conversion, strconv for strings

### Types
- boolean
- numeric
  - integer
  - floating point
  - complex number
- text type

Constants
### Arrays and Slices
__arrays__
- Collection of items of the same Type
- Fixed size
- Declare:
  - a:= [3]int{1,2,3}
  - a:= [...]int {1,2,3}
  - var a [3]int
- len() return size

__slices__
- Backed by array
- Declare
  - var a []int
  - a = make([]int, 3, 5)
  - slice existing
- len() return length
- cap() return capacity
- append() add elements, may trigger copy
- _Copies refer to  the same underlying array_

### Maps and Structs
__Map__
- created via literal or make()
- check for presence using value, ok, if not found
  ok = false, value = zero value
- copies by reference

__Struct__
- collection of disparate data types
- field names follow the capitalization == export to global rules
- anonymous structs
- copies by value, like arrays, can be key to maps
- no inheritance, can use embedding as "has-a"
- tags used to describe field

### Defer, Panic, Recover
__defer__ 
- execute command after the last line of the function, before returning
    - used to delay execution of a statement until function exits
- in LIFO order, thus useful for closing resources
    - used to group "open" and "close" functions together
- input parameters of the function are evaluated at time of defer call, eagerly, not at the end of the function

__panic__
- defer statements will execute before the panic executes! So it goes defer -> panic -> function return!
- you can recover from a panic in the deferred function

### Pointers
- the `new` syntax
- syntactic sugar means pointer.field === (*pointer).foo. * de-reference has lower precedence 
  than .
  
### Functions
- entry point: package main, main function.
- function inside functions actually do have access to the local variable declared in the 
  outer scope. BUT, don't use this, use parameters to pass in for stability in async.
- methods
  - value receiver
  - pointer receiver

### Interface
- Polymorphism! 
  Also in the reverse way! You can interface to match the type struct.
- Convention: if only one method, then name is `method-er`. e.g. `reader`.
- You can compose interface together to build the interface you want.
- Cast interface to type structs, also supports "result, ok" syntax for checking if the 
  cast to type struct succeeds
  ```
    var wc = NewBufferedWriterCloser()
    bwc, ok := wc.(*BufferedWriterCloser)
  ```
- empty interface is the wildcard of type : `interface{}`. Used as intermediate step before 
  we type convert it into useful interfaces or types. 
  ```
    var v interface{}
    switch (v.(type)){...}
  ```
- IMPORTANT ABOUT POINTERS. 
  When you implement an interface with the value of the type. I.e:
  ```
    var obj Interface = ImplementingType{}
  ```
  Your method set is only the methods of the interface with __value receivers__.
  But when you implement the interface like this:
  ```
    var obj Interface = &ImplementingType{}
  ```
  
  Then your method set is both methods with __pointer receivers__ and __value receivers__.
  
- BEST PRACTICES
  - use many, small interfaces
    - `io.Writer, io.Reader, interface{}`
  - don't export interfaces for types that will be consumed
  - do export interfaces for types that will be used by package
  - design functions and methods to receive interfaces whenever possible
  

### Concurrent programming. Go Routines
- most language use OS Kernel threads with 1MiB of overhead, then you be conservative to use these expensive threads.
  In go, use an abstraction of a thread. Inside the go runtime, scheduler maps goroutiue to low-level threads,
  for periods of time. Scheduler than assign goroutines to threads on that time. 
  Goroutine are assign small stack spaces, because they can be reallocated quickly. 
  Therefore, it's common to see 10K+ goroutines running at the same time.
- Go has closures, and anonymous function has the closure of outside functions. so we can do:
```
  var msg = "hello"
  go func(){
    fmt.Println(msg)
  }()
```
  - decouple with parameters
    ```
      var msg = "hello"
      go func(){
        fmt.Println(msg)
      }(msg)
    ```
- use `WaitGroup` and `RWMutex` to synchronize green threads
- `runtime.GOMAXPROCS(num)` sets the number of threads available to the runtime. 
  Often it runs faster when more than one thread per core. You TUNE it.
- Best practice:
  - don't create goroutines in libraries. let user control concurrency
  - when creating goroutine, know how it will end. 
    (watcher goroutines, avoid subtle memory leaks)
  - __IMPORTANT__ check race conditions at compile time.
    - go compiler add `-race` flag to detect data races.

### Channels
- The reason you choose Go. It stands out from other language, 
  which were designed with the single-processor model.
- Channel Basics
    - create: `make(chan <type>)`
    - by default: unbuffered channels. when pushing to channel, it blocks the thread 
      until some thread pulls the message out.
      so if no thread pull messages out, it's a deadlock.
- Restricting Data Flow
    - usually only want goroutines to have only one of reader/writer role to channels
- Buffered Channels
  - `make(chan int, 50)` makes buffered channels, so we can push more than one message
    before any pulls.
  - buffered channels ACTUALLY block the receiver until message is available. 
    NOTICE HOW IT'S THE EXACT OPPOSITE OF NORMAL CHANNELS. 
- Close Channels
  - goes hand in hand with range loops. The range loops wants to read all the messages from 
    the channel, so it stops only when the channel is closed. If go routines exit without
    closing the channel, the range loop read thread deadlocks, since its wait for a channel
    that will never be closed.
  - don't push messages to a closed channel. You CAN'T EVEN detect a channel is closed 
    unless you try to push and try to recover from a panic.
- For range loops w/ Channels
  - `for val := range ch {...}` gets the value
  - The under-the-hood way is:
    ```
      for {
        if val, ok := <- ch; ok {
          // do something with val
        } else {
          break
        }
      }
    ```
- Select w/ Channels
  - default, just like a blocking select in C.
  - Common practice to use with a doneChannel, which is used purely to send a signal.
    `stuct{}` actually cost no memory allocation.
    ```
      type logEntry struct {...}
      main thread:
      logCh := make(chan logEntry, 50)
      doneCh := make(chan struct{})
      
      logCh <- logEntry{...}
      ...
      doneCh <- struct{}{} // define and initialize struct with no fields
    
      logger thread:
      select {
        case entry := <-logCh:
          // do something from logCh
        case entry := <-doneCh:
          // clean up
          break
        // having a default code block makes the select a non-blocking select in C.
        // defaut: ... 
      }
    ```
  
