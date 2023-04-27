package main

import "fmt"

func main() {
	//As in many other language, you can use 'continue', 'break' within for loops

	//CONTINUE skips the iteration
	for i := 0; i < 10; i++ {
		if i%2 != 0 { //is odd
			fmt.Println("skip iteration")
			continue
		}
		// this will not execute if i is odd
		fmt.Printf("i is even : %v \n", i)
	}

	//BREAK will stop the whole loop and resume at the next line right after the loop
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration %v \n", i)
		if i == 2 {
			fmt.Println("break")
			break
		}
	}
	fmt.Println("next line")

	//RETURN will behave like a normal return and stop the function entirely
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration %v \n", i)
		if i == 2 {
			fmt.Println("Return")
			return
		}
	}
	fmt.Println("This will never execute !")
}
