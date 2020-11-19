package structure

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
		ss := Structure{testCase.InputSl}
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
		ss := Structure{testCase.InputSl}
		ss.Delete(testCase.InputVal)

		if !isSliceEquals(ss.items, testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, ss.items)
		}
	}
}

func TestUpdate(t *testing.T) {
	ss := Structure{}
	input := []int{10, 4, 3, 9, -5, -3}
	expected := []int{4, 9, 10}

	for _, v := range input {
		ss.Update(v)
	}

	if !isSliceEquals(ss.items, expected) {
		t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", expected, ss.items)
	}
}

func TestGetMin(t *testing.T) {
	st := Structure{[]int{5, 12, 33}}

	if st.GetMin() != 5 {
		t.Fatalf("Function GetMin return not min value")
	}
}

func TestGetMax(t *testing.T) {
	st := Structure{[]int{5, 12, 33}}

	if st.GetMax() != 33 {
		t.Fatalf("Function GetMax return not max value")
	}
}

func BenchmarkInsertIntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 1) }
func BenchmarkInsertIntoMiddle(b *testing.B)       { benchmarkInsert(b, 10000, 5000) }
func BenchmarkInsertIntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 10000) }
func BenchmarkInsertBatch(b *testing.B)            { benchmarkInsertBatch(b, 10000) }

func BenchmarkDeleteFromBeginning(b *testing.B) { benchmarkDelete(b, 10000, 0) }
func BenchmarkDeleteFromMiddle(b *testing.B)    { benchmarkDelete(b, 10000, 5000) }
func BenchmarkDeleteFromEnd(b *testing.B)       { benchmarkDelete(b, 10000, 9999) }

func BenchmarkGetMinOn100(b *testing.B)   { benchmarkGetMin(b, 100) }
func BenchmarkGetMinOn10000(b *testing.B) { benchmarkGetMin(b, 10000) }

func BenchmarkGetMaxOn100(b *testing.B)   { benchmarkGetMax(b, 100) }
func BenchmarkGetMaxOn10000(b *testing.B) { benchmarkGetMax(b, 10000) }

func benchmarkInsert(b *testing.B, originalSize, value int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := Structure{getTestSet(originalSize)}
		originalLen := len(ss.items)
		b.StartTimer()

		ss.Insert(value)

		if originalLen >= len(ss.items) {
			b.Fatal()
		}
	}
}

func benchmarkInsertBatch(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := Structure{getTestSet(size / 2)}
		b.StartTimer()

		for j := 0; j < size; j++ {
			ss.Insert(j)
		}
	}
}

func benchmarkDelete(b *testing.B, originalSize, value int) {
	items := getTestSet(originalSize)
	tmpSlice := make([]int, len(items))

	for i := 0; i < b.N; i++ {
		copy(tmpSlice, items)
		ss := Structure{tmpSlice}

		ss.Delete(value)

		if len(items) <= len(ss.items) {
			b.Fail()
		}
	}
}

func benchmarkGetMin(b *testing.B, size int) {
	st := Structure{getTestSet(size)}

	for i := 0; i < b.N; i++ {
		if st.GetMin() != 0 {
			b.Fail()
		}
	}
}

func benchmarkGetMax(b *testing.B, size int) {
	st := Structure{getTestSet(size)}

	for i := 0; i < b.N; i++ {
		if st.GetMax() != size-1 {
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
