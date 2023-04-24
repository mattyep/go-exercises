package main

import "fmt"

/*
OBJECTIVE :
The goal here is to display all the possible zero-values
of initialised (but unassigned) variables depending on their types
*/
func main() {
	//ACTION :
	// 1 - Run the following code
	// 3 - Notice that initialized variables are not all Nil by default depending on their types
	// 4 - Notice how some types are aliases for others (e.g. byte, rune)

	var myInt int
	var myInt32 uint32
	var myFloat32 float32
	var myByte byte
	var myString string
	var myRune rune
	var myArray [3]int
	var mySlice []int
	var myBoolean bool
	var myPointer *int

	fmt.Printf("The 'zero-value' of a %T : %v \n", myInt, myInt)
	fmt.Printf("The 'zero-value' of a %T : %v \n", myInt32, myInt32)
	fmt.Printf("The 'zero-value' of a %T : %f \n", myFloat32, myFloat32)
	fmt.Printf("The 'zero-value' of a %T : %v \n", myByte, myByte)
	fmt.Printf("The 'zero-value' of a %T : '%s' \n", myString, myString)
	fmt.Printf("The 'zero-value' of a %T : %v \n", myRune, myRune)
	fmt.Printf("The 'zero-value' of a %T : %v \n", myArray, myArray)
	fmt.Printf("The 'zero-value' of a %T : %v \n", mySlice, mySlice)
	fmt.Printf("Is the unassigned slice %v comparable to nil ? %v \n", mySlice, mySlice == nil)
	fmt.Printf("The 'zero-value' of a %T : %v \n", myBoolean, myBoolean)
	fmt.Printf("The 'zero-value' of a pointer %T : %v \n", myPointer, myPointer)
}
