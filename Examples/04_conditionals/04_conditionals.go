package main

import (
	"errors"
	"fmt"
)

func main() {
	// ➡️ if statements are farely simple. Just note that you don't need the parenthesis
	value := "foo"
	if value == "foo" {
		//foo
	} else if value == "bar" {
		//bar
	} else {
		//else
	}

	// ➡️ You can also insert assignement's statements within the if
	if result := (1 + 1); result == 2 {
		//the scope of the variable 'result' is within the if/else statement only
	} else if result%2 == 0 {
		//'result' also exists here
	} else {
		//and also here
		fmt.Println(result)
	}
	//However, the 'result' variable no longer exists here

	// ➡️ This syntax is particulary useful when you need to get the result
	// of a function ONLY to test it with a If / Switch straight after.
	//e.g.
	if err := writeInDatabase(); err != nil {
		fmt.Println(err)
	}
	//the variable 'err' no longer exists here
	//You can redeclare it with the same name after
	var err error
	fmt.Println(err)
}

func writeInDatabase() error {
	// Write to database and eventually produce errors
	return errors.New("Failed to write in Database : Database is down")
}
