package main

import (
	"reflect"
	"testing"
)

// 1 - Complete the following method with a classic for-loop so that
// the array returned contains the values 1 to 5 ([1 2 3 4 5])
func arrayWith1to5() (array [5]int) {
	// ...
	return array
}

// 2 - Complete the following method with a for-loop of your choice so that
// the array returned contains the values 1 to n ([1 2 ... n])
func arrayWith1toN(n int) []int {
	var array []int = make([]int, n)
	// ...
	return array
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
