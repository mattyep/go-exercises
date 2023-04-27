package main

import "fmt"

// Package-level constant
const Pi = 3.14159265359

func main() {
	// Function-level constant
	const localConstant = "Foo"

	fmt.Println(Pi, localConstant)

	// ✅ Constants are immutable variable, you can't modify them !
	// ⛔ localConstant = "Bar" -> cannot assign to localConstant (untyped string constant "Foo")
}
