package main

import (
	"reflect"
	"testing"
)

// 1️⃣ - Complete the following method with a classic for-loop so that
// the returned array contains the values 1 to 5 ([1 2 3 4 5])
// hint 1 : 'var array [5]int' defines a fixed size array which results into [0, 0, 0, 0, 0]
// You can already iterate on it to modify it's values
// hint 2 : The index variable you get from the for-loop can help you BOTH modify the value at the given index AND
// also generate the expected values (1 to 5) since they are just shifted by one from the index
func arrayWith1to5() [5]int {
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i + 1
	}
	return array
}

// 2️⃣ - Complete the following method with a for-RANGE loop over 'array' in order to
// fill it with values going from 1 to ... n (ex: [1, 2, 3, 4, 5, ...., n])
// hint 1: make() prefills the array with (n) zero-values. Therefore the 'array' variable has already the expected size.
// hint 2: You will have to use the index given by the for-loop to BOTH modify the array's values AND
// generate the values 1 .. n which are just shifted by one compared to the index values.
func arrayWith1toN(n int) []int {
	var array []int = make([]int, n)
	for i := range array {
		array[i] = i + 1
	}
	return array
}

// 3️⃣ - Complete the following method with a for-loop to do the sum of all
// the values contained in the param 'array', then return this sum.
// The array can have different sizes.
// ex A: array: [1, 2, 3] -> return 6
// ex B: array: [2, 2, 2, 2, 2] - > return 10
// Hint: You will need a 'sum' variable outside the loop where you can 'accumulate' each of the values
// you treat on the loop
func arraySum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

// 4️⃣ - Complete the following method with a for-loop of your choice so that
// it squares all the values (value * value) and return them on an array of the same size.
// ex A: array: [1, 2, 3] -> return [1, 4, 9]
// ex B: array: [2, 2, 2, 2, 2] - > return [4, 4, 4, 4, 4]
func arraySquare(array []int) []int {
	var result []int = make([]int, len(array))
	for i, value := range array {
		result[i] = value * value
	}
	return result
}

// 5️⃣ - Complete the following method with a for-loop of your choice so that
// all the values are shifted to the left and the first goes at the end. (e.g. [1 2 3 4] -> [2 3 4 1])
// see : https://www.javatpoint.com/program-to-left-rotate-the-elements-of-an-array
// The goal is to play with for-loop and index + 1. Using 'append(array[1:], array[0])' here is cheating :).
func arrayRotateLeft(array []int) []int {
	var result []int = make([]int, len(array))
	first := array[0]

	for i := 0; i < len(array)-1; i++ {
		result[i] = array[i+1]
	}

	result[len(array)-1] = first
	return result
}

// ###################################################################
// TEST

func TestArrayWith1to5(t *testing.T) {
	expected := [5]int{1, 2, 3, 4, 5}
	if arr := arrayWith1to5(); !reflect.DeepEqual(arr, expected) {
		t.Errorf("arrayWith1to5 returned %v instead of %v", arr, expected)
	}
}

func TestArrayWith1toN(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	if arr := arrayWith1toN(5); !reflect.DeepEqual(arr, expected) {
		t.Errorf("arrayWith1toN returned %v instead of %v", arr, expected)
	}
}

func TestSumArray(t *testing.T) {
	testSum := []int{1, 2, 3}
	if sum := arraySum(testSum); sum != 6 {
		t.Errorf("sumArray returned %v instead of %v", sum, 6)
	}
}

func TestArraySquare(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 4, 9}
	if arr := arraySquare(input); !reflect.DeepEqual(arr, expected) {
		t.Errorf("arraySqaure returned %v instead of %v", arr, expected)
	}
}

func TestArrayRotateLeft(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{2, 3, 4, 1}
	if arr := arrayRotateLeft(input); !reflect.DeepEqual(arr, expected) {
		t.Errorf("arrayRotateLeft returned %v instead of %v", arr, expected)
	}
}
