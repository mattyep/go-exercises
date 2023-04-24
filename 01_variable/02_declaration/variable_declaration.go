package main

/*
OBJECTIVE :
Understand the difference between the 'var' and the ':=' syntax
*/
func main() {
	// ########################################################################
	// 1 - define a variable (a) of type int without any value

	// display your variable with fmt.Printf("a: %T(%v) \n", a, a)

	// ########################################################################
	// 2 - define a variable (b) with a value 10 (preferably with the := syntax)

	// You didn't have to declare the type because the compiler infered it from the value (Integer).
	// (!) Use the 'var' syntax when you don't have any value to assign yet and the zero-value is good enough (e.g int -> 0)
	// (!) Use the ':=' syntax when you declare a new variable directly with a value (e.g. 10) which is not the default zero-value

	// ########################################################################
	// However, the default type might not be the one you want. So :
	// 3 - Define a variable (c) of type uint64 with the value 10

	// 4 - Now, do the same same thing with a variable (d) but with the ':=' syntax. You will have to cast the value into a uint64 (e.g. 'unint64(10)')

	// 3 and 4 do the same things but the 'var' syntax is preferable here.
	// (!) Use the 'var' syntax when you want a type which is not the default one.

	// ########################################################################
	// 5 - Use the fooBar() method below and assign its result into two variable 'foo' and 'bar'
	// (see https://go.dev/doc/effective_go#redeclaration)

	// 6 - Now, replace the var 'foo' with an underscore _ .
	// This way, you explicitly say you're not interested by one of the values of fooBar and you are not forced to use it.
}

func fooBar() (string, string) {
	return "foo", "bar"
}
