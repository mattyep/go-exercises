package main

import "fmt"

func main() {
	//There is 4 ways to do for-loops

	//#########################################################
	// 1️⃣ : FOR-RANGE

	//First, let's start with the most popular and used syntax
	myArray := []int{1, 2, 3}
	for index, value := range myArray {
		// This syntax iterate over an array and give you the index and the value for each iteration
		// 📢 the 'value' variable is a copy. If you change it, it will not impact the array.
		// e.g. value += 1 has no impact
		fmt.Printf("For-range loop : index %v, value %v \n", index, value)

		//You can use the index to access the array directly :
		myArray[index] += 1
	}
	//By naming convention you will mostly see this syntax used with shorter variable names
	//You can use whatever name which suit your code. the index is usually i but the value variable name can be contextualized.
	for i, v := range myArray {
		fmt.Printf("For-range loop 2 : index %v, value %v \n", i, v)
	}

	//If you are not interested by the index and you only want the value, you can ignore the index with an _
	for _, v := range myArray {
		fmt.Printf("For-range loop value-only : value %v \n", v)
	}

	//If you only want the index and not the value, simply don't assign it
	for i := range myArray {
		fmt.Printf("For-range loop index-only : index %v \n", i)
	}

	//#########################################################
	// 2️⃣ : FOR-CLASSIC

	//Second way is a classical for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("For-classic loop with i : iteration %v \n", i)
	}

	//that, btw, you can adapt as you which
	for foo := 5; foo < 10; foo += 2 {
		fmt.Printf("For-classic loop tuned with i : iteration %v \n", foo)
	}
	//backward
	for i := 10; i >= 0; i-- {
		fmt.Printf("Final Countdown ... %v \n", i)
	}

	//#########################################################
	// 3️⃣ : FOR-CONDITIONAL

	//Here is the third form of for loop in golang
	//It's the equivalent to 'while' in other languages
	counter := 0
	for counter < 3 {
		fmt.Printf("For-conditional loop with counter : iteration %v \n", counter)
		counter++
	}

	//#########################################################
	// 4️⃣ : FOR-INFINITE

	// Forth way is the infinite loop
	counter2 := 0
	for {
		fmt.Printf("For-infinite loop with counter2 : iteration %v \n", counter)
		counter2++
		if counter2 >= 3 {
			//⚠️ Make sure you always have an exit condition !
			break
		}
	}

}
