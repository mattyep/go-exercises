package main

import "fmt"

// FizzBuzz
// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print "Fizz" instead of the number
// and for the multiples of five print "Buzz".
// For numbers which are multiples of both three and five print "FizzBuzz".
// For more details : https://codingdojo.org/kata/FizzBuzz/
func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 && i != 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i%3 == 0 && i != 0 {
			fmt.Printf("Fizz\n")
		} else if i%5 == 0 && i != 0 {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
