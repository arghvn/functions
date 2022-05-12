package main

import "fmt"

// Here's a function that takes two `int`s and returns their sum as an 'int'

func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't automatically return the value of the last expression.

	return a + b

}

// When you have multiple consecutive parameters of the same type,
//you may omit the type name for the like-typed parameters up to the final parameter that declares the type.

func plusPlus(a, b, c int) int {

	return a + b + c
}

//Call a function just as you’d expect, with name(args).

//There are several other features to Go functions. One is multiple return values.

func main() {

	res := plus(1, 2)

	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)

	fmt.Println("1+2+3 =", res)

}

// output :
// 1+2 = 3
// 1+2+3 = 6

// more details :
//To define a simple function, pay attention to the implementation pattern of the plus function.
// The input parameters must always be specified. In this example, we pass two parameters a and b to the function by typing int.
//After determining the input parameters and before inserting the code inside the bracket for the function,
// the keyword int is placed, which means that the output value of the function must be data int.
//If the data type of all input parameters is the same, the data type keyword can be mentioned at the end of all parameters,
// like the plusPlus function. To call a function, just use the args (name) syntax.
