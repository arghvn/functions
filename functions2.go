//Introduce another method of defining a function.

package main

import "fmt"

func main() {

plus := func(a int, b int) int {

return a + b

}

res := plus(1, 2)

fmt.Println("1+2 =", res)

//output :
//1+2=3
