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
}
/*
Output:

Hello GeeksforGeeks
string
Happy 3.14 Day
Go rules? true
*/
