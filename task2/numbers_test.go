package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		InputSl  []int
		InputVal int
		Expected []int
	}{
		{
			InputSl:  []int{},
			InputVal: 1,
			Expected: []int{1},
		},
		{
			InputSl:  []int{1},
			InputVal: 2,
			Expected: []int{1, 2},
		},
		{
			InputSl:  []int{2},
			InputVal: 1,
			Expected: []int{1, 2},
		},
		{
			InputSl:  []int{1, 3},
			InputVal: 2,
			Expected: []int{1, 2, 3},
		},
		{
			InputSl:  []int{1, 3, 9},
			InputVal: 7,
			Expected: []int{1, 3, 7, 9},
		},
		{
			InputSl:  []int{1, 3, 9},
			InputVal: 3,
			Expected: []int{1, 3, 3, 9},
		},
		{
			InputSl:  []int{1, 3, 9},
			InputVal: -3,
			Expected: []int{-3, 1, 3, 9},
		},
	}

	for _, testCase := range testCases {
		res := Insert(testCase.InputSl, testCase.InputVal)

		if !isSliceEquals(res, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, res)
		}
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		InputSl  []int
		InputVal int
		Expected []int
	}{
		{
			InputSl:  []int{},
			InputVal: 1,
			Expected: []int{},
		},
		{
			InputSl:  []int{1},
			InputVal: 2,
			Expected: []int{1},
		},
		{
			InputSl:  []int{1, 2},
			InputVal: 1,
			Expected: []int{2},
		},
		{
			InputSl:  []int{1, 1, 2},
			InputVal: 1,
			Expected: []int{1, 2},
		},
		{
			InputSl:  []int{-1, 2},
			InputVal: -1,
			Expected: []int{2},
		},
	}

	for _, testCase := range testCases {
		res := Delete(testCase.InputSl, testCase.InputVal)

		if !isSliceEquals(res, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, res)
		}
	}
}

func TestUpdateNumbers(t *testing.T) {
	testCases := []struct {
		Numbers      []int
		InputNumbers []int
		Expected     []int
	}{
		{
			Numbers:      []int{},
			InputNumbers: []int{10, 4, 3, 9, -5, -3},
			Expected:     []int{4, 9, 10},
		},
		{
			Numbers:      []int{1, 5, 6},
			InputNumbers: []int{11, 3, 2, -6},
			Expected:     []int{1, 2, 3, 5, 11},
		},
	}

	for _, testCase := range testCases {
		for _, v := range testCase.InputNumbers {
			testCase.Numbers = UpdateNumbers(testCase.Numbers, v)
		}

		if !isSliceEquals(testCase.Numbers, testCase.Expected) {
			t.Fatalf(
				"Failed compare that two slices are equals. Want %v, got %v",
				testCase.Expected,
				testCase.Numbers,
			)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	testSet := getTestSet()

	for i := 0; i < b.N; i++ {
		res := Insert(testSet, i)

		if len(testSet) >= len(res) {
			b.Fatal()
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	testSet := getTestSet()

	for i := 0; i < b.N; i++ {
		// Will include an insert execution. Find b.StopTimer and b.StartTimer methods, but with them benchmark hangs.
		res := Delete(Insert(testSet, i), i)

		if len(testSet) != len(res) {
			b.Fatal()
		}
	}
}

func isSliceEquals(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for i, v := range sl1 {
		if v != sl2[i] {
			return false
		}
	}

	return true
}

func getTestSet() []int {
	testSet := make([]int, 0, 1000)

	for i := 1; i < 1000; i++ {
		testSet = append(testSet, i)
	}

	return testSet
}
