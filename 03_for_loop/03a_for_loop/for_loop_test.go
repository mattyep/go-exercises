package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 1️⃣ - Complete the following method with a classic for-loop so that
// the returned array contains the values 1 to 5 ([1 2 3 4 5])

func arrayWith1to5() [5]int {
	var array = [5]int{1, 2, 3, 4, 5}

	for _, j := range array {
		fmt.Println(j)
	}
	return array
}

func arrayWith1to5bis() [5]int {
	var array [5]int
	//array -> [0 0 0 0 0]
	for i := 0; i <= 4; i++ {
		array[i] = i + 1
	}
	return array
}

// 2️⃣ - Complete the following method with a for-loop of your choice so that
// the returned array contains the values 1 to n ([1 2 ... n])
func arrayWith1toN(n int) []int {

	var slice []int
	// slice -> [] -> nil
	slice = make([]int, n)
	// slice -> [0 0 0 ... 0] n times
	// len(array) == n

	//  for index := 0; index < n; index++ {
	// 		slice[index] += 1 + index
	//  	//slice = append(slice, index + 1)
	// 	fmt.Println(array)
	//  }
	for index, valeur := range slice {
		slice[index] = (1 + index)
		fmt.Println(valeur)
	}
	return slice

}

// 3️⃣ - Complete the following method with a for-loop of your choice so that
// the sum of the array is returned (e.g. [1 2 3] -> 6)
func arraySum(array []int) int {

	//var tabArray = []int{1, 2, 3}

	sum := 0

	for index, valeur := range array {
		sum += valeur
		fmt.Println(index, valeur)

	}
	// for index := 0; index <= 2; index++ {
	// 	fmt.Println(sum)
	// }

	return sum

}

// 4️⃣ - Complete the following method with a for-loop of your choice so that
// it squares all the values (value * value) and return them on an array of the same size (e.g. [1 2 3] -> [1 4 9])
func arraySquare(array []int) []int {
	var result = []int{1, 2, 3}

	for index, valeur := range array {

		fmt.Println(result)
	}
	// return result
}

// 5️⃣ - Complete the following method with a for-loop of your choice so that
// all the values are shifted to the left and the first goes at the end. (e.g. [1 2 3 4] -> [2 3 4 1])
// see : https://www.javatpoint.com/program-to-left-rotate-the-elements-of-an-array
// The goal is to play with for-loop and index + 1. Using 'append(array[1:], array[0])' here is cheating :).
func arrayRotateLeft(array []int) []int {
	var result []int = make([]int, len(array))
	// ...
	return result
}

// ###################################################################
// TEST

func TestArrayWith1to5(t *testing.T) {
	expected := [5]int{1, 2, 3, 4, 5}
	if arr := arrayWith1to5bis(); !reflect.DeepEqual(arr, expected) {
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
		t.Errorf("arraySquare returned %v instead of %v", arr, expected)
	}
}

func TestArrayRotateLeft(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{2, 3, 4, 1}
	if arr := arrayRotateLeft(input); !reflect.DeepEqual(arr, expected) {
		t.Errorf("arrayRotateLeft returned %v instead of %v", arr, expected)
	}
}
