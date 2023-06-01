package main

import "fmt"

// FizzBuzz again, but with a switch case now :)
// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print "Fizz" instead of the number
// and for the multiples of five print "Buzz".
// For numbers which are multiples of both three and five print "FizzBuzz".
// For more details : https://codingdojo.org/kata/FizzBuzz/
func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0 && i != 0:
			fmt.Printf("FizzBuzz\n")
		case i%3 == 0 && i != 0:
			fmt.Printf("Fizz\n")
		case i%5 == 0 && i != 0:
			fmt.Printf("Buzz\n")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
