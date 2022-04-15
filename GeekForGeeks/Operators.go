// Go program to illustrate the
// use of bitwise operators
package main

import "fmt"

func main() {
p:= 34
q:= 20

// & (bitwise AND)
result1:= p & q
fmt.Printf("Result of p & q = %d", result1)

// | (bitwise OR)
result2:= p | q
fmt.Printf("\nResult of p | q = %d", result2)

// ^ (bitwise XOR)
result3:= p ^ q
fmt.Printf("\nResult of p ^ q = %d", result3)

// << (left shift)
result4:= p << 1
fmt.Printf("\nResult of p << 1 = %d", result4)

// >> (right shift)
result5:= p >> 1
fmt.Printf("\nResult of p >> 1 = %d", result5)

// &^ (AND NOT)
result6:= p &^ q
fmt.Printf("\nResult of p &^ q = %d \n", result6)

a := 4

// Using address of operator(&) and
// pointer indirection(*) operator
b := &a
fmt.Println(&a)
fmt.Println(*b)
*b = 7
fmt.Println(a)


}

/*
OUtput:

Result of p & q = 0
Result of p | q = 54
Result of p ^ q = 54
Result of p << 1 = 68
Result of p >> 1 = 17
Result of p &^ q = 34
0xc0000aa058
4
7
*/
