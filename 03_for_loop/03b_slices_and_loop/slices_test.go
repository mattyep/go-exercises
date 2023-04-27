package main

import (
	"reflect"
	"testing"
)

// What are slices ?? 🍰🤔
// It's just arrays that are dynamically sized !
// [5]int -> array (fixed size of 5)
// []int -> slice (no fixed size)
// Not only that "arrays" have a fixed size but this size must be defined at compile time !
//
// ✅ You will rarely be able to define such size during development time, therefore,
// what you will manipulate, almost exclusively, are called Slices.
// (And yes, arrays_test.go was already about slices, oups..)

// 1️⃣ - Complete the addElementTo method to return
// the slice with one more value (whatever integer you want), using the append method.
// hint: the append() method returns a new slice with the extra value.
// e.g. addElementTo([]int{}) -> [1]
func addElementTo(slice []int) []int {
	//...
	return nil
}

// 2️⃣ - Complete the addMultipleElementTo method to return
// the slice with (nbElement) more values, using the append method and a for-loop.
// The value itself doesn't matter.
// e.g. addMultipleElementTo([]int{}, 5) -> [1 1 1 1 1]
func addMultipleElementTo(slice []int, nbElement int) []int {
	//...
	return nil
}

// 3️⃣ - Complete the makeSlice method to return
// a new slice provisioned with (size) elements. Use the make() method to do it.
// e.g. makeSlice(5) -> [0 0 0 0 0]
func makeSlice(size int) []int {
	//...
	return nil
}

// 4️⃣ - Complete the copySlice method to return
// a copy of the slice. You will have to use the make() and copy() method. Google it :)
func copySlice(slice []int) []int {
	//...
	return nil
}

// 5️⃣ - Complete the following method to iterate on the given slice and
// return the first odd number (e.g. 1 3 5 ...) you find in it.
// Hint : An odd number can be tested with '(value % 2) != 0'
func findOddNumberInSlice(slice []int) int {
	//...
	return 0
}

// ###################################################################
// TEST

func TestAddElementTo(t *testing.T) {
	if s := addElementTo([]int{}); len(s) != 1 {
		t.Errorf("addElementTo() should have appended 1 element but added %v", len(s))
	}
}

func TestAddMultipleElementTo(t *testing.T) {
	if s := addMultipleElementTo([]int{}, 5); len(s) != 5 {
		t.Errorf("addMultipleElementTo([]int{}, 5) should have appended 5 elements but added %v", len(s))
	}
}

func TestMakeSlice(t *testing.T) {
	if s := makeSlice(5); len(s) != 5 {
		t.Errorf("makeSlice(5) should have returned a slice with 5 elements but had %v elements", len(s))
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3}
	if copy := copySlice(original); !reflect.DeepEqual(original, copy) {
		t.Errorf("copySlice() should have returned a copy of %v but returned %v instead", original, copy)
	}
	if copy := copySlice(original); &original[0] == &copy[0] {
		t.Errorf("copySlice() should have returned a copy of %v but returned the same slice instead", original)
	}
}

func TestFindOddNumberInSlice(t *testing.T) {
	testArray := []int{0, 2, 4, 5}
	if r := findOddNumberInSlice(testArray); r != 5 {
		t.Errorf("findOddNumberInSlice() returned %v instead of 5 from %v", r, testArray)
	}
}
