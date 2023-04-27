package main

import (
	"fmt"
)

/*
OBJECTIVE :
Understand the difference between the 'var' and the ':=' syntax
*/
func main() {
	// ########################################################################
	// 1️⃣ - define a variable (a) of type int without any value
	// 1 - define a variable (a) of type int without any value
	var a int
	fmt.Println(a)
	// display your variable with fmt.Printf("a: %T(%v) \n", a, a)

	// ########################################################################
	// 2 - define a variable (b) with a value 10 (preferably with the := syntax)
	b := 10
	fmt.Println(b)
	// You didn't have to declare the type because the compiler infered it from the value (Integer).
	// ✅ Use the 'var' syntax when you don't have any value to assign yet and the zero-value is good enough (e.g int -> 0)
	// ✅ Use the ':=' syntax when you declare a new variable directly with a value (e.g. 10) which is not the default zero-value

	// ########################################################################
	// However, the default type might not be the one you want. So :
	// 3 - Define a variable (c) of type uint64 with the value 10

	//...

	d := uint64(10)
	fmt.Printf("d: %T(%v) \n", d, d)
	// 3 and 4 do the same things but the 'var' syntax is preferable here.
	// ✅ Use the 'var' syntax when you want a specific type which is not the default one.

	// ########################################################################
	// 5️⃣ - Call the fooBar() method and assign its result into two variable varA and varB
	// (see https://go.dev/doc/effective_go#redeclaration)

	fooBar()

	mickey, minnie := fooBar()
	fmt.Println(mickey, minnie)
	// 6️⃣ - Now, replace the var 'varA' with an underscore _ .
	// This way, you explicitly say you're not interested by one of the values of fooBar and you are not forced to use it.

	_, minnie = fooBar()
	fmt.Println(minnie)

}

func fooBar() (string, string) {
	return "foo", "bar"
}
