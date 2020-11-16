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
		{[]int{}, 1, []int{1}},
		{[]int{1}, 2, []int{1, 2}},
		{[]int{2}, 1, []int{1, 2}},
		{[]int{1, 3}, 2, []int{1, 2, 3}},
		{[]int{1, 3, 9}, 7, []int{1, 3, 7, 9}},
		{[]int{1, 3, 9}, 3, []int{1, 3, 3, 9}},
		{[]int{1, 3, 9}, -3, []int{-3, 1, 3, 9}},
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
		{[]int{}, 1, []int{}},
		{[]int{1}, 2, []int{1}},
		{[]int{1, 2}, 1, []int{2}},
		{[]int{1, 1, 2}, 1, []int{1, 2}},
		{[]int{-1, 2}, -1, []int{2}},
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
		{[]int{}, []int{10, 4, 3, 9, -5, -3}, []int{4, 9, 10}},
		{[]int{1, 5, 6}, []int{11, 3, 2, -6}, []int{1, 2, 3, 5, 11}},
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

// The worst scenario in single call but in batch call a little bit better then V2.
func BenchmarkInsertV1IntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 0, 1) }
func BenchmarkInsertV1IntoMiddle(b *testing.B)       { benchmarkInsert(b, 10000, 500, 1) }
func BenchmarkInsertV1IntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 999, 1) }
func BenchmarkInsertV1Batch(b *testing.B)            { benchmarkInsertBatch(b, 10000, 1) }

// Good scenario in single call, but show bad performance in batch call against V3 and even V1.
func BenchmarkInsertV2IntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 0, 2) }
func BenchmarkInsertV2IntoMiddle(b *testing.B)       { benchmarkInsert(b, 10000, 500, 2) }
func BenchmarkInsertV2IntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 999, 2) }
func BenchmarkInsertV2Batch(b *testing.B)            { benchmarkInsertBatch(b, 10000, 2) }

// In single call lose to V2 but in batch call perfect.
func BenchmarkInsertV3IntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 0, 3) }
func BenchmarkInsertV3IntoMiddle(b *testing.B)       { benchmarkInsert(b, 10000, 500, 3) }
func BenchmarkInsertV3IntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 999, 3) }
func BenchmarkInsertV3Batch(b *testing.B)            { benchmarkInsertBatch(b, 10000, 3) }

// Maybe its good enough.
func BenchmarkDeleteFromTheBeginning(b *testing.B) { benchmarkDelete(b, 10000, 0) }
func BenchmarkDeleteFromMiddle(b *testing.B)       { benchmarkDelete(b, 10000, 500) }
func BenchmarkDeleteFromTheEnd(b *testing.B)       { benchmarkDelete(b, 10000, 999) }

func benchmarkInsert(b *testing.B, size, index, ver int) {
	testSet := getTestSet(size)
	value := testSet[index]

	for i := 0; i < b.N; i++ {
		var res []int

		switch ver {
		case 2:
			res = InsertV2(testSet, value)
		case 3:
			res = InsertV3(testSet, value)
		default:
			res = Insert(testSet, value)
		}

		if len(testSet) >= len(res) {
			b.Fatal()
		}
	}
}

func benchmarkInsertBatch(b *testing.B, size, ver int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testSet := getTestSet(size / 2)
		b.StartTimer()

		for j := 0; j < size; j++ {
			switch ver {
			case 2:
				testSet = InsertV2(testSet, j)
			case 3:
				testSet = InsertV3(testSet, j)
			default:
				testSet = Insert(testSet, j)
			}
		}
	}
}

func benchmarkDelete(b *testing.B, size, index int) {
	testSet := getTestSet(size)
	testSetCopy := make([]int, len(testSet))
	value := testSet[index]

	for i := 0; i < b.N; i++ {
		copy(testSetCopy, testSet)
		res := Delete(testSetCopy, value)

		if len(testSet) <= len(res) {
			b.Fail()
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

func getTestSet(size int) []int {
	testSet := make([]int, 0, size)

	for i := 0; i < size; i++ {
		testSet = append(testSet, i)
	}

	return testSet
}
