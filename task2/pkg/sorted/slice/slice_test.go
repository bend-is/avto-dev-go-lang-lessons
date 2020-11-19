package slice

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
		res := Insert(testCase.InputSl, testCase.InputVal)

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
		res := Delete(testCase.InputSl, testCase.InputVal)

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
			testCase.Numbers = Update(testCase.Numbers, v)
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
//func BenchmarkInsertV1IntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 0, Insert) }
//func BenchmarkInsertV1IntoMiddle(b *testing.B)    { benchmarkInsert(b, 10000, 500, Insert) }
//func BenchmarkInsertV1IntoTheEnd(b *testing.B)      { benchmarkInsert(b, 10000, 999, Insert) }
//func BenchmarkInsertV1Batch(b *testing.B)     { benchmarkInsertBatch(b, 10000, Insert) }

// Good scenario in single call, but show bad performance in batch call against V3 and even V1.
//func BenchmarkInsertV2IntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 0, InsertV2) }
//func BenchmarkInsertV2IntoMiddle(b *testing.B)    { benchmarkInsert(b, 10000, 500, InsertV2) }
//func BenchmarkInsertV2IntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 999, InsertV2) }
//func BenchmarkInsertV2Batch(b *testing.B)     { benchmarkInsertBatch(b, 10000, InsertV2) }

// In single call lose to V2 but in batch call perfect.
func BenchmarkInsertIntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 0, InsertV3) }
func BenchmarkInsertIntoMiddle(b *testing.B)       { benchmarkInsert(b, 10000, 500, InsertV3) }
func BenchmarkInsertIntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 999, InsertV3) }
func BenchmarkInsertBatch(b *testing.B)            { benchmarkInsertBatch(b, 10000, InsertV3) }

// Maybe its good enough.
func BenchmarkDeleteFromBeginning(b *testing.B) { benchmarkDelete(b, 10000, 0) }
func BenchmarkDeleteFromMiddle(b *testing.B)    { benchmarkDelete(b, 10000, 500) }
func BenchmarkDeleteFromEnd(b *testing.B)       { benchmarkDelete(b, 10000, 999) }

func benchmarkInsert(b *testing.B, size, index int, insertFunc func([]int, int) []int) {
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

func benchmarkInsertBatch(b *testing.B, size int, insertFunc func([]int, int) []int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testSet := getTestSet(size / 2)
		b.StartTimer()

		for j := 0; j < size; j++ {
			testSet = insertFunc(testSet, j)
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
