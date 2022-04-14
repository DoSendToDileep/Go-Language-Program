/*
Data types specify the type of data that a valid Go variable can hold.
In Go language, the type is divided into four categories which are as follows:

Basic type: Numbers, strings, and booleans come under this category.
Aggregate type: Array and structs come under this category.
Reference type: Pointers, slices, maps, functions, and channels come under this category.
Interface type

*/
// Go program to illustrate
// the use of floating-point
// numbers
package main
import "fmt"

func main() {
	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b-a

	// Display the result
	fmt.Printf("Result is: %f", c)

	// Display the type of c variable
	fmt.Printf("\nThe type of c is : %T \n", c)

  var aa complex128 = complex(6, 2)
var bb complex64 = complex(9, 2)
fmt.Println(aa)
fmt.Println(bb)

// Display the type
fmt.Printf("The type of a is %T and "+
         "the type of b is %T", aa, bb)

       // Multiple variables of different types
       // are declared and initialized in the single line
       var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56
}

/*
Output:
Result is: 14.440000
The type of c is : float64
(6+2i)
(9+2i)
The type of a is complex128 and the type of b is complex64
[Finished in 3.291s]

*/
