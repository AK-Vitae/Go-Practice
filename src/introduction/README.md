# Introductory Notes

**Go (golang)**

* Strong and statically typed
* Excellent community
* Key features:
  * Simplicity
  * Fast compile times
  * Garbage collected
  * Built-in concurrency
  * Compile to standalone binaries

## Variables

### Variable Declaration:

* ```go
  var foo int
  var foo int = 32
  foo := 42
  ```

* Can't redeclare variables, but can shadow them
* All variables must be used

### Visibility

* lower case first letter for package scope
* upper case first letter to export
* no private scope

### Type conversions

* ```go
  destinationType(variable)
  ```

* use strconv package for strings

## Primitive Types

### Boolean type

* Values are true or false
* Not an alias for other types
* Zero value is false

### Numeric Types

* **Integers**
  * Signed Integers
    * int type has a varying size, but min 32 bits
    * 8bit(int8) through 64 bit (int64)
  * Unsigned integers
    * 8 bit(byte and uint8) through 32 bit(uint32)
  * Arithmetic operations: addition, subtraction, multiplication, division, and remainder
  * Bitwise operations: and, or, xor, and not
  * Zero value is 0
* **Floating point numbers**
  * Literal Styles:
    * Decimal(3.14)
    * Exponential(13e18 or 2E10)
    * Mixed(13.7e12)
  * Arithmetic operations: addition, subtraction, multiplication, and division
  * Zero value is 0
* **Complex numbers**
  * Built-in functions
    * complex - make complex number from two floats
    * real - get real part as float
    * imag - get imaginary part as float
  * Arithmetic operations: addition, subtraction, multiplication, and division
  * Zero value is 0+0i

### Text types

* **Strings**
  * UTF-8
  * Immutable
  * Can be concatenated with plus(+) operator
  * Can be converted to []byte
* **Rune**
  * UTF-32
  * Alias for int32
  * Special methods normally required to process
    * e.g. strings.Reader#ReadRune

## Constants

* Immutable but can be shadowed
* Replaced by the compiler at compile time
  * Value must be calculable at compile time
* Typed constants work like immutable variables
* Untyped constants work like literals

### Enumerated Constants

* Special symbol **iota** allows related constants to be created easily
* Iota starts at 0 in each const block and increments by one
* Watch out of constant values that match zero values for variables

### Enumerated expressions

* Operations that can be determined at compile time are allowed
  * Arithmetic 
  * Bitwise operations
  * Bitshifting

## Arrays

* Collection of items with same type

* Fixed size

* Declaration styles:
  
  * ```go
    a := [3]int{1, 2, 3}
    a := [...]int{1, 2, 3}
    var a [3]int
    ```
  
* Access via zero-based index

* **len** function returns size of array

* Copies refer to **different underlying data**

## Slices

* Backed by array
* Creation styles:
  * Slice existing array or slice
  
  * Literal style
  
  * Via **make** function:
    
    * ```go
      a := make([]int, 10) // create slice with capacity and length == 10
      a := make([]int, 10, 100) // create slice with length == 10 and capacity == 100
      ```
* **len** function returns length of slice
* **cap** function returns length of underlying array
* **append** function to add elements to slice
* Copies refer to **same underlying data**

## Maps

* Collections of value types that are accessed via keys
* Created via literals or via **make** function
* Members accessed via [key] syntax
  * myMap["key"] = "value"
* Check for presence with "value, ok" form of result
* Multiple assignments refer to **same underlying data**

## Structs

* Collections of disparate data types that describe a single concept
* Keyed by named fields
* Normally created as types, but anonymous structs are allowed
* Structs are value types
* No inheritance, but can use composition via embedding
* Tags can be added to struct fields to describe fields

## Control Flow

### If statements

* Can be initalized via a code block
* Comparison Operators <, >, <=,>=, !=
* Using error parameters is good practice

### Switch Statements

* You can switch on a tag or be tag less
* Can test multiple cases but they have to be unique
* Breaks are implied
* **fallthrough** keyword needs to be explicitly stated
* Type Switch: Can be helpful when transforming data

### Looping

* simple loops:
  * for initializer; test; incrementor {}
  * for test {}
  * for {}
* exiting early
  * **break**
  * **continue**
  * **labels**
* looping over collections
  * arrays, slices, maps, strings, and channels
  * for k, v := range collection {}

### Defer

* Used to delay execution of a statement until function exits
* Useful to group "open" and "close" functions together (Need to be careful in loops)
* Run in LIFO (last in, first out) order 
* Arguments evaluated at time defer is executed, not at time of called function execution

### Panic

* Occur when program cannot continue at all 
* Use for unrecoverable events but don't use when a file can't be opened, unless it is critical
* Order of execution: functions -> deferred statements -> panic statements -> return value

### Recover

* Used to recover from panic and only useful in deferred functions
* Current functions will not attempt to continue, but higher function in call stack will

## Pointers

### Creating Pointers

* Pointer types use an asterisk(*) as a prefix to type pointed to
  * *int - a pointer to an integer
* Use the **addressof** operator(&) to get address of variable

### Dereferencing Pointers

* Dereference a pointer by preceding with an asterisk(*)
* Complex types (e.g. structs) are automatically dereferenced

### Create pointers to objects

* Can use the addressof operator(&) if value types already exists

  * ```go
    ms := myStruct{foo: 42}
    p := &ms
    ```

* Use addressof operator before initializer

  * ```go
    &myStruct{foo: 42}
    ```

* Use the **new** keyword

  * Can't initialize fields at the same time

### Types with internal pointers

* All assignments operations in Go are copy operations
* Slices and maps contain internal pointers, so copies point to same underlying data

## Functions

### Basic Syntax

```go
func foo() {
	...
}
```

### Parameters

* Comma delimited list of variables and types

  * ```go
    func foo(bar string, baz int)
    ```

* Parameters of same type list type once

  * ```go
    func foo(bar, baz int)
    ```

* When pointers are passed in, the functions can change the value in the caller

  * This is always true for data of slices and maps

* Use variadic parameters to send list of same types in

  * Must be last parameter

  * Received as a slice

  * ```go
    func foo(bar string, baz ...int)
    ```

### Return Values

* Single return values: just list type
* Multiple return value: list types surrounded by parentheses
  * the (result type, error) paradigm is a very common idiom
* Can use named return values
* Can return addresses of local variables

### Anonymous Functions

* Functions don't have names if they are:

  * Immediately invoked

    * ```go
      func() {
          ...
      }()
      ```

  * Assigned to a variable or passed as an argument to a function 

    * ```go
      a := func() {
          ...
      }
      a()
      ```

### Functions as types

* Can assign functions to variables or use as arguments and return values in functions

* Type signature is like function signature, with no parameter names

  * ```go
    var f func(string, string, int) (int, error)
    ```

### Methods

* Functions that executes in context of a type

* Format

  * ```go
    func (g greeter) greet() {
    	...
    }
    // g is a value receiver and gets copy of type
    func (g *greeter) greet() {
    	...
    }
    // g is a pointer receiver and gets pointer to type
    ```