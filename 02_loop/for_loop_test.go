package main

import (
	"reflect"
	"testing"
)

// 1️⃣ - Complete the following method with a classic for-loop so that
// the returned array contains the values 1 to 5 ([1 2 3 4 5])
func arrayWith1to5() [5]int {
	var array [5]int
	// ...
	return array
}

// 2️⃣ - Complete the following method with a for-loop of your choice so that
// the returned array contains the values 1 to n ([1 2 ... n])
func arrayWith1toN(n int) []int {
	var array []int = make([]int, n)
	// ...
	return array
}

// 3️⃣ - Complete the following method with a for-loop of your choice so that
// the sum of the array is returned (e.g. [1 2 3] -> 6)
func arraySum(array []int) int {
	//...
	return 0
}

// 4️⃣ - Complete the following method with a for-loop of your choice so that
// it squares all the values (value * value) and return them on an array of the same size (e.g. [1 2 3] -> [1 4 9])
func arraySquare(array []int) []int {
	var result []int
	//...
	return result
}

// 5 - Complete the following method with a for-loop of your choice so that
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
