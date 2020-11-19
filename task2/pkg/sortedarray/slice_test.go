package sortedarray

import "testing"

func TestInsertToSlice(t *testing.T) {
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
		res := InsertToSlice(testCase.InputSl, testCase.InputVal)

		if !isSliceEquals(res, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, res)
		}
	}
}

func TestDeleteFromSlice(t *testing.T) {
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
		res := DeleteFromSlice(testCase.InputSl, testCase.InputVal)

		if !isSliceEquals(res, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, res)
		}
	}
}

func TestUpdateSlice(t *testing.T) {
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
			testCase.Numbers = UpdateSlice(testCase.Numbers, v)
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
//func BenchmarkSliceInsertV1Beginning(b *testing.B) { benchmarkSliceInsert(b, 10000, 0, InsertToSlice) }
//func BenchmarkSliceInsertV1Middle(b *testing.B)    { benchmarkSliceInsert(b, 10000, 500, InsertToSlice) }
//func BenchmarkSliceInsertV1eEnd(b *testing.B)      { benchmarkSliceInsert(b, 10000, 999, InsertToSlice) }
//func BenchmarkSliceInsertV1Batch(b *testing.B)     { benchmarkSliceInsertBatch(b, 10000, InsertToSlice) }

// Good scenario in single call, but show bad performance in batch call against V3 and even V1.
//func BenchmarkSliceInsertV2Beginning(b *testing.B) { benchmarkSliceInsert(b, 10000, 0, InsertToSliceV2) }
//func BenchmarkSliceInsertV2Middle(b *testing.B)    { benchmarkSliceInsert(b, 10000, 500, InsertToSliceV2) }
//func BenchmarkSliceInsertV2End(b *testing.B)       { benchmarkSliceInsert(b, 10000, 999, InsertToSliceV2) }
//func BenchmarkSliceInsertV2Batch(b *testing.B)     { benchmarkSliceInsertBatch(b, 10000, InsertToSliceV2) }

// In single call lose to V2 but in batch call perfect.
func BenchmarkSliceInsertBeginning(b *testing.B) { benchmarkSliceInsert(b, 10000, 0, InsertToSliceV3) }
func BenchmarkSliceInsertMiddle(b *testing.B)    { benchmarkSliceInsert(b, 10000, 500, InsertToSliceV3) }
func BenchmarkSliceInsertEnd(b *testing.B)       { benchmarkSliceInsert(b, 10000, 999, InsertToSliceV3) }
func BenchmarkSliceInsertBatch(b *testing.B)     { benchmarkSliceInsertBatch(b, 10000, InsertToSliceV3) }

// Maybe its good enough.
func BenchmarkSliceDeleteFromBeginning(b *testing.B) { benchmarkDeleteFromSlice(b, 10000, 0) }
func BenchmarkSliceDeleteFromMiddle(b *testing.B)    { benchmarkDeleteFromSlice(b, 10000, 500) }
func BenchmarkSliceDeleteFromEnd(b *testing.B)       { benchmarkDeleteFromSlice(b, 10000, 999) }

func benchmarkSliceInsert(b *testing.B, size, index int, insertFunc func([]int, int) []int) {
	testSet := getTestSet(size)
	value := testSet[index]

	for i := 0; i < b.N; i++ {
		var res []int

		res = insertFunc(testSet, value)

		if len(testSet) >= len(res) {
			b.Fatal()
		}
	}
}

func benchmarkSliceInsertBatch(b *testing.B, size int, insertFunc func([]int, int) []int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testSet := getTestSet(size / 2)
		b.StartTimer()

		for j := 0; j < size; j++ {
			testSet = insertFunc(testSet, j)
		}
	}
}

func benchmarkDeleteFromSlice(b *testing.B, size, index int) {
	testSet := getTestSet(size)
	testSetCopy := make([]int, len(testSet))
	value := testSet[index]

	for i := 0; i < b.N; i++ {
		copy(testSetCopy, testSet)
		res := DeleteFromSlice(testSetCopy, value)

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
