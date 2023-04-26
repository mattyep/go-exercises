package main

import (
	"fmt"
	"testing"
)

// 1️⃣ - Complete the firstItemOf method to return
// the first element of the input array.
func firstItemOf(array []int) int {
	//fmt.Println(array[0])
	return array[0]
}

// 2️⃣ - Complete the findItemByIndex method to return
// the array's element positioned at the given index.
func findItemByIndex(array []int, index int) int {
	// array []int & index are values
	// tout ce quon met entre [] sera un index
	return array[index]
}

// 3️⃣ - Complete the lengthOf method to return
// the array's length.
func lengthOf(array []int) int {
	//fmt.Println("la taille de mon tableau1 :", [len(array)])
	return len(array)
}

// 4️⃣ - Complete the lastItemOf method to return
// the last element of the input array. (hint: use the len() method)
func lastItemOf(array []int) int {
	fmt.Println("Dernier item =", array[len(array)-1])
	return array[len(array)-1]
}

// ###################################################################
// TEST

type arrayTest struct {
	array []int
	first int
	last  int
}

var arrayTests = []arrayTest{
	{[]int{1, 2, 3}, 1, 3},
	{[]int{3, 2, 1, 0}, 3, 0},
	{[]int{2, 6, 1, 0, 3}, 2, 3},
}

func TestFirstItemOf(t *testing.T) {
	for _, test := range arrayTests {
		if i := firstItemOf(test.array); i != test.first {
			t.Errorf("firstItemOf returned %d instead of %d", i, test.first)
		}
	}
}

func TestFindItemByIndex(t *testing.T) {
	for _, test := range arrayTests {
		if i := findItemByIndex(test.array, test.array[0]); i != test.array[test.array[0]] {
			t.Errorf("firstItemOf returned %d instead of %d", i, test.array[test.array[0]])
		}
	}
}

func TestLengthOf(t *testing.T) {
	for _, test := range arrayTests {
		if i := lengthOf(test.array); i != len(test.array) {
			t.Errorf("firstItemOf returned %d instead of %d", i, len(test.array))
		}
	}
}

func TestLastItemOf(t *testing.T) {
	for _, test := range arrayTests {
		if i := lastItemOf(test.array); i != test.last {
			t.Errorf("firstItemOf returned %d instead of %d", i, test.last)
		}
	}
}
