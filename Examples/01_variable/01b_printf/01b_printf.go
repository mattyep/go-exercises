package main

import "fmt"

func main() {
	//SEE all options here : https://pkg.go.dev/fmt

	var myInt int = 0x0001F60E
	var myFloat float64 = 3.14159265359

	// Use %v by default, Printf will adapt
	fmt.Printf("Display any value with %%v : %v \n", myFloat)
	// Use %T if you want the type of the value
	fmt.Printf("Display the type of a value with %%T (upper-case) : %T \n", myInt)
	// Use %d if you want to display an integer in base-10
	fmt.Printf("Display a base-10 decimal value with %%d : %d \n", myInt)
	// Use %X if you want to display an integer in base-16 (usually memory adresses)
	fmt.Printf("Display a base-16 decimal value with %%X : %X \n", myInt)
	// Use %c if you want the want the character represented by a Unicode code point
	fmt.Printf("Display the Unicode code point with %%c : %c \n", myInt)

	//✅ Just use %v and for the rest, go see https://pkg.go.dev/fmt
}
