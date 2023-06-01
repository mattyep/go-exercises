package main

import "testing"

// 1️⃣ - Complete the firstItemOf method to return
// the first element of the input array.
func firstItemOf(array []int) int {
	return array[0]
}

// 2️⃣ - Complete the findItemByIndex method to return
// the array's element positioned at the given index.
// hint: Same as above 1️⃣, except that it's not 0 but the 'index' param.
func findItemByIndex(array []int, index int) int {
	return array[index]
}

// 3️⃣ - Complete the lengthOf method to return
// the array's length.
func lengthOf(array []int) int {
	return len(array)
}

// 4️⃣ - Complete the lastItemOf method to return
// the last element of the input array. (hint: use the len() method)
func lastItemOf(array []int) int {
	return array[len(array)-1]
}

// 5️⃣ - Exact same exercise as number 2️⃣. Except now the index might be out of bound
// and generate a Panic. Complete the safeFindItemByIndex method to return
// the array's element positioned at the given index OR 0 if the given index is out of bound
// Hint: You will have first to compare the given index param against the length of the array
// in order to make sure that the index does exist.
func safeFindItemByIndex(array []int, index int) int {
	if index < 0 || index >= len(array) {
		return 0
	}
	return array[index]
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

func TestSafeFindItemByIndex(t *testing.T) {
	for _, test := range arrayTests {
		if i := safeFindItemByIndex(test.array, test.array[0]); i != test.array[test.array[0]] {
			t.Errorf("firstItemOf returned %d instead of %d", i, test.array[test.array[0]])
		}
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("safeFindItemByIndex raised a panic : %v", r)
				}
			}()
			if i := safeFindItemByIndex(test.array, len(test.array)); i != 0 {
				t.Errorf("safeFindItemByIndex called with out of bound index %v returned %d instead of %d", len(test.array), i, test.array[test.array[0]])
			}
		}()
	}
}
