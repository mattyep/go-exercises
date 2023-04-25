package main

import "fmt"

func main() {
	//Basic var declaration
	// a is declared without any value. The zero-value of int is 0
	var a int
	fmt.Println(a)

	//Let's assign a value to (a)
	a = 10
	fmt.Println(a)

	//We can also use a simpler syntax
	b := 10
	fmt.Println(b)
	//It only works if we have a value to assign upfront.
	// We don't have to declare the type because the compiler infered it from the value (10 is an integer).
	// ✅ Use the 'var' syntax when you don't have any value to assign yet and the zero-value is good enough (e.g int -> 0)
	// ✅ Use the ':=' syntax when you declare a new variable directly with a value (e.g. 10) which is not the default zero-value

	// ⛔ DON'T DO THAT :
	var c int = 0 //valid code but useless assignement.
	fmt.Println(c)

	//other examples :
	foo := "bar"
	// hardcoded floats are represented like 1.0
	float := 1.0

	// You can define similar variables at once
	var x, y int
	// you can define AND assign multiple variables (of any type) at once
	i, j := 1, "toto"

	fmt.Println(foo, float, x, y, i, j)
}
