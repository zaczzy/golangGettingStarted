## Notes
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