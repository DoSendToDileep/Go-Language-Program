package main
import "fmt"
const PI = 3.14
func main(){
	const GFG = "GeeksforGeeks"
	fmt.Println("Hello", GFG)
  fmt.Printf("%T \n",GFG)

  //GFG = "ASADF"
  //Error: cannot assign to GFG (untyped string constant "GeeksforGeeks")
	fmt.Println("Happy", PI, "Day")

	const Correct= true
	fmt.Println("Go rules?", Correct)

  const untypedInteger          = 123
const untypedFloating    = 123.12

const typedInteger  int             = 123
const typedFloatingPoint   float64  = 123.12

fmt.Printf("%T %T %T %T \n",untypedInteger,untypedFloating,typedInteger,typedFloatingPoint)
fmt.Println(untypedInteger,untypedFloating,typedInteger,typedFloatingPoint)
}

/*
Output:

Hello GeeksforGeeks
string
Happy 3.14 Day
Go rules? true
int float64 int float64
123 123.12 123 123.12
*/
