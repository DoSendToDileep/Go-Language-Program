package main
import "fmt"
func main() {

 var name = "Chakravarthi"
 fmt.Println("Hello, Dileep ", name)

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int64
    fmt.Println(e)

    /*
    This will give error as it overflows..
    var ee int64 = 9223372036854775808
    */

    f := `apple`
    fmt.Println(f)

    coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
    fmt.Println(coral)
}
/*
Output:
Hello, Dileep  Chakravarthi
initial
1 2
true
0
apple
[blue coral staghorn coral pillar coral]

=================
Other relevant information
=================
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     IEEE-754 32-bit floating-point numbers
float64     IEEE-754 64-bit floating-point numbers
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

uint     unsigned, either 32 or 64 bits
int      signed, either 32 or 64 bits
uintptr  unsigned integer large enough to store the uninterpreted bits of a pointer value

*/


/*
Rules for Defining Identifiers: There are certain valid rules for defining a valid Go identifier. These rules should be followed, otherwise, we will get a compile-time error.

The name of the identifier must begin with a letter or an underscore(_). And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.
The name of the identifier should not start with a digit.
The name of the identifier is case sensitive.
Keywords is not allowed to use as an identifier name.
There is no limit on the length of the name of the identifier, but it is advisable to use an optimum length of 4 – 15 letters only.

*/
