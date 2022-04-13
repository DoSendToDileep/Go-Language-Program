/*
Keywords or Reserved words are the words in a language that are used for some internal process or represent some predefined actions.
These words are therefore not allowed to use as an identifier. Doing this will result in a compile-time error.
*/
package main
import "fmt"

// Here, package, import, func,
// var are keywords
func main() {

// Here, a is a valid identifier
var a = "GeeksforGeeks"

fmt.Println(a)
// Here, the default is an
// illegal identifier and
// compiler will throw an error
// var default = "GFG"
}

/*
I. Declaration
  The keywords are used to declare all kinds of code elements in Go.
1.	const
  The ‘const’ keyword is used to introduce a name for a scalar value like 2 or 3.14, etc.
2.	var
  The ‘var’ keyword is used to create the variables in the ‘Go’ language.
3.	func
  The ‘func’ keyword is used to declare a function.
4.	type
  We can use the ‘type’ keyword to introduce a new struct type.
5.	import
  The ‘import’ keyword is used to import packages.
6.	package
  The code is grouped as a unit in a package. The ‘package’ keyword is used to define one.

II. Composite Types
The keywords are used to indicate some composite type.
1.	chan
  The ‘chan’ keyword is used to define a channel. In ‘Go’, you are allowed to run parallel pieces of code simultaneously.
2.	interface
  The ‘interface’ keyword is used to specify a method set. A method set is a list of methods for a type.
3.	map
  The ‘map’ keyword defines a map type. A map os an unordered collection of the key-value pairs.
4.	struct
  Struct is a collection of fields. We can use the ‘struct’ keyword followed by the field declarations.


III. Control Flow
The keywords are used to control flow of code.
1.	break
  The ‘break’ keyword lets you terminate a loop and continue execution of the rest of the code.
2.	case
  This is a form of a ‘switch’ construct. We specify a variable after the switch.
3.	continue
  The ‘continue’ keyword allows you to go back to the beginning of the ‘for’ loop.
4.	default
  The ‘default’ statement is optional but you need to use either case or default within the switch statement. The switch jumps to the default value if the case does not match the controlling expression.
5.	else
  The ‘else’ keyword is used with the ‘if’ keyword to execute an alternate statement if a certain condition is false.
6.	fallthrough
  This keyword is used in a case within a switch statement. When we use this keyword, the next case is entered.
7.	for
  The ‘for’ keyword is used to begin a for loop.
8.	goto
  The ‘goto’ keyword offers a jump to a labeled statement without any condition.
9.	if
  The ‘if’ statement is used to check a certain condition within a loop.
10.	range
  The ‘range’ keyword lets you iterate items over the list items like a map or an array.
11.	return
  Go allows you to use the return values as variables and you can use the ‘return’ keyword for that purpose.
12.	select
  The ‘select’ keyword lets a goroutine to wait during the simultaneous communication operations.
13.	switch
  The ‘switch’ statement is used to start a loop and use the if-else logic within the block.


IV. Function Modifier
The keywords also control flow keywords, but in other specific manners. They modify function calls.
1.	defer
  The ‘defer’ keyword is used to defer the execution of a function until the surrounding function executes.
2.	go
  The ‘go’ keyword triggers a goroutine which is managed by the golang-runtime.



*/
