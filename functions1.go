# Example of function with 1 parameter and no returned value
package main
import (
 "fmt"
)
// A function with one parameter with that does not return a value
func SayHi(name string) {
 fmt.Printf("Hi %s\n", name)
}
// A function with two parameters that does return a value (int)
func AddTwoNumbers(a int, b int) int {
 return a + b
}
// A function with two parameters (both int) that return also two
// values (int)
func Division(c, d int) (int, int) {
 quotient := c / d
 reminder := c % d
 return quotient, reminder
}
func main() {
 SayHi("Paulina")
 fmt.Println(AddTwoNumbers(3, 9))
 q, r := Division(16, 5)
 fmt.Println("Division between", 16, " and 5 , is: Quotient ", q, " reminder ", r)
}
