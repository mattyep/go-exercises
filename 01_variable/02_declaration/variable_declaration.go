package main

/*
OBJECTIVE :
Understand the difference between the 'var' and the ':=' syntax
*/
func main() {
	// ########################################################################
	// 1️⃣ - define a variable (a) of type int without any value
	// display your variable with fmt.Printf("a: %T(%v) \n", a, a)

	//...

	// ########################################################################
	// 2️⃣ - define a variable (b) with a value 10 (preferably with the := syntax)

	//...

	// You didn't have to declare the type because the compiler infered it from the value (Integer).
	// ✅ Use the 'var' syntax when you don't have any value to assign yet and the zero-value is good enough (e.g int -> 0)
	// ✅ Use the ':=' syntax when you declare a new variable directly with a value (e.g. 10) which is not the default zero-value

	// ########################################################################
	// However, there is one rare exception to the above rules.
	// When you use the := syntax, the default type decided by the compiler might not be
	// the one you want. For instance, you might want an uint64 instead of a basic int.
	// Let's try this :
	// 3️⃣ - Define a variable (c) of type uint64 with the value 1_000_000_000 with the var syntax.

	//...

	// 4️⃣ - Now, do the same same thing with a variable (d) but with the ':=' syntax.
	// You will have to cast the value 1_000_000_000 into a uint64 before assigning it (Search engines are your friends)

	//...

	// 3 and 4 do the same things but the 'var' syntax is preferable here.
	// ✅ Use the 'var' syntax when you want a specific type which is not the default one.

	// ########################################################################
	// 5️⃣ - Use the fooBar() method below and assign its result into two variable 'foo' and 'bar'
	// (see https://go.dev/doc/effective_go#redeclaration)

	//...

	// 6️⃣ - Now, replace the var 'foo' with an underscore _ .
	// This way, you explicitly say you're not interested by one of the values of fooBar and you are not forced to use it.

	//...
}

func fooBar() (string, string) {
	return "foo", "bar"
}
