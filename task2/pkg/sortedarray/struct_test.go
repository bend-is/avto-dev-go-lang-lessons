package sortedarray

import "testing"

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
		{[]int{1, 3, 9}, 3, []int{1, 3, 3, 9}},
		{[]int{1, 3, 9}, -3, []int{-3, 1, 3, 9}},
	}

	for _, testCase := range testCases {
		ss := SortedStruct{testCase.InputSl}
		ss.Insert(testCase.InputVal)

		if !isSliceEquals(ss.items, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, ss.items)
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
		ss := SortedStruct{testCase.InputSl}
		ss.Delete(testCase.InputVal)

		if !isSliceEquals(ss.items, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, ss.items)
		}
	}
}

func TestUpdateNumbers(t *testing.T) {
	ss := SortedStruct{}
	input := []int{10, 4, 3, 9, -5, -3}
	expected := []int{4, 9, 10}

	for _, v := range input {
		ss.Update(v)
	}

	if !isSliceEquals(ss.items, expected) {
		t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", expected, ss.items)
	}
}

func BenchmarkStructInsertBeginning(b *testing.B) { benchmarkStructInsert(b, 10000, 1) }
func BenchmarkStructInsertMiddle(b *testing.B)    { benchmarkStructInsert(b, 10000, 5000) }
func BenchmarkStructInsertEnd(b *testing.B)       { benchmarkStructInsert(b, 10000, 10000) }
func BenchmarkStructInsertBatch(b *testing.B)     { benchmarkStructInsertBatch(b, 10000) }

func BenchmarkStructDeleteBeginning(b *testing.B) { benchmarkStructDelete(b, 10000, 0) }
func BenchmarkStructDeleteMiddle(b *testing.B)    { benchmarkStructDelete(b, 10000, 5000) }
func BenchmarkStructDeleteEnd(b *testing.B)       { benchmarkStructDelete(b, 10000, 9999) }

func benchmarkStructInsert(b *testing.B, originalSize, value int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := SortedStruct{getTestSet(originalSize)}
		originalLen := len(ss.items)
		b.StartTimer()

		ss.Insert(value)

		if originalLen >= len(ss.items) {
			b.Fatal()
		}
	}
}

func benchmarkStructInsertBatch(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := SortedStruct{getTestSet(size / 2)}
		b.StartTimer()

		for j := 0; j < size; j++ {
			ss.Insert(j)
		}
	}
}

func benchmarkStructDelete(b *testing.B, originalSize, value int) {
	items := getTestSet(originalSize)
	tmpSlice := make([]int, len(items))

	for i := 0; i < b.N; i++ {
		copy(tmpSlice, items)
		ss := SortedStruct{tmpSlice}

		ss.Delete(value)

		if len(items) <= len(ss.items) {
			b.Fail()
		}
	}
}
