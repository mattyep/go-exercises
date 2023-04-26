package main

import "fmt"

// What are slices ?? 🍰🤔
// It's just arrays that are dynamically sized !
// [5]int -> array (fixed size of 5)
// []int -> slice (no fixed size)
// Not only that "arrays" have a fixed size but this size must be defined at compile time !
//
// ✅ You will rarely be able to define such size during development time, therefore,
// what you will manipulate, almost exclusively, are called Slices.
// (And yes, arrays_test.go was already about slices, oups..)

func main() {

	// ➡️ A slice is always define with another type
	var sliceOfString []string
	// An array is the same thing but defined with a static size
	var arrayOfString [3]string

	//-----------------------------------------------------------------------------------
	// ➡️ an empty slice is comparable to nil.
	//It's important to note because, by contrast, a fixed-size array would never be nil
	// e.g.
	//    'var s []int' -> [] -> nil
	//    'var a [3]int' -> [0 0 0]
	if sliceOfString == nil {
		fmt.Printf("👉🏼 Empty slice %v is comparable to nil ! \n", sliceOfString)
	}
	fmt.Printf("👉🏼 In contrast, arrays are never empty %v, they are filled with zero-values (here, empty strings \"\") ! \n", arrayOfString)

	//-----------------------------------------------------------------------------------
	// ➡️ Array or slices, you can access their values with whatever
	// means you want which results to an int
	myArray := [5]int{4, 7, 2, 8, 1}
	value := myArray[4] // with a literal integer
	lastIndex := 4
	value = myArray[lastIndex]      // with an integer variable
	value = myArray[lastIndex-1]    //with an arithmetic operation resulting to an int
	value = myArray[len(myArray)-1] // with call to a function resulting to an int
	value = myArray[func() int {
		return 12 / 3
	}()] // with an anonymous function, seriously ?? Yes, WHATEVER that results into an int.

	fmt.Printf("👉🏼 You can use any means you want to compute the index you want :  \n", value)

	//-----------------------------------------------------------------------------------
	// ⚠️ Be careful with index manipulation, they should stay within the array/slice actual size.
	outOfBoundIndex := 5 //myArray's last index is 4, 5 doesn't exist.
	func() int {
		defer func() { fmt.Println("😱 PANIC : ", recover()) }()
		return myArray[outOfBoundIndex] //panic: runtime error: index out of range [5] with length 5
	}()
	//Here's the way to stay safe :
	if outOfBoundIndex < len(myArray) {
		//👍 Keep calm and check array's boundaries
	}

	//-----------------------------------------------------------------------------------
	// ➡️ For Slices only, use append() to add a new element dynamically
	sliceOfString = append(sliceOfString, "foo")
	sliceOfString = append(sliceOfString, "bar")

	fmt.Println("👉🏼 sliceOfString after 2 append() : ", sliceOfString) //-> ["foo", "bar"]

	//-----------------------------------------------------------------------------------
	// ➡️ You can also create a new slice and fill it with zero-values
	sliceOfInt := make([]int, 3) //-> [0 0 0]
	//Doing so is almost equivalent to
	var arrayOfInt [3]int //-> [0 0 0]
	//Except that you can define the size of the slice dynamically at runtime
	//while the array needs its size hardcoded at development time
	fmt.Println("👉🏼 Static vs dynamic size, same result at the end :", sliceOfInt, "<->", arrayOfInt)

	//-----------------------------------------------------------------------------------
	// ✅ Slices bear this name 🍰 because you can create them from subset of other arrays (or other slices).
	// ➡️ You can extract part of an array/slice content with the [:] syntax
	firstHalf := sliceOfString[:1]  //slice from beginning to index 1 : [foo]
	secondHalf := sliceOfString[1:] //slice from index 1 to end : [bar]
	fmt.Printf("👉🏼 sliceOfString[:1] equals to %v while sliceOfString[1:] equals to %v \n", firstHalf, secondHalf)

	//☢️ Be very careful with the [:] because the values you extract are still referenced to the original Slice.
	// If you modify one of this value, it will also change the original Slice.
	// e.g.
	firstHalf[0] = "foo-modified"
	fmt.Println("👉🏼 firstHalf modified sliceOfString: ", sliceOfString) //-> ["foo-modified", "bar"]
	//And even worse !! 👇
	firstHalf = append(firstHalf, "bar-overflowed!!")
	fmt.Println("👉🏼 firstHalf also modified secondHalf 😱 : ", secondHalf)    //-> ["bar-overflowed!!"]
	fmt.Println("👉🏼 and it modified sliceOfString (again): ", sliceOfString) //-> ["foo-modified", "bar-overflowed!!"]
	// ✅ CONCLUSION : Slices are easily mutable and modification should be treated carefully

	//-----------------------------------------------------------------------------------
	// ➡️ However, you can make a real copy of the slice or array in order to modify it without risk
	// You have to provision a new slice with the same size first
	sliceOfStringCopy := make([]string, len(sliceOfString))
	copy(sliceOfStringCopy, sliceOfString)

	sliceOfStringCopy[0] = "foo-copy"
	//did it modify the original sliceOfString ?
	fmt.Println("👉🏼 No, the copy did not modify the original sliceOfString: ", sliceOfString) //-> ["foo-modified", "bar-overwritten!!"]

	//-----------------------------------------------------------------------------------

	//-----------------------------------------------------------------------------------
	// ➡️ Last but not least, strings behaves almost like slice !!
	fact := "String behaves like a slice"
	factLength := len(fact)
	factWord := fact[:6]
	fmt.Println("👉🏼 ", factLength, factWord)
}
