package slice

import (
	"errors"
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
		{[]int{1, 3, 9}, 3, []int{1, 3, 3, 9}},
		{[]int{1, 3, 9}, -3, []int{-3, 1, 3, 9}},
	}

	for _, testCase := range testCases {
		ss := Slice{testCase.InputSl}
		ss.Insert(testCase.InputVal)

		if !isSliceEquals(ss.GetItems(), testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, ss.GetItems())
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
		{[]int{1, 1, 1}, 1, []int{1, 1}},
		{[]int{-1, 2}, -1, []int{2}},
	}

	for _, testCase := range testCases {
		ss := Slice{testCase.InputSl}
		ss.Delete(testCase.InputVal)

		if !isSliceEquals(ss.GetItems(), testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, ss.GetItems())
		}
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		Input    []int
		Expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 1, -1}, []int{1}},
		{[]int{-1, -1, -1}, []int{}},
		{[]int{10, 4, 3, 9, -5, -3}, []int{4, 9, 10}},
	}

	for _, testCase := range testCases {
		ss := Slice{}

		ss.Update(testCase.Input...)

		if !isSliceEquals(ss.GetItems(), testCase.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", testCase.Expected, ss.GetItems())
		}
	}
}

func TestGetMin(t *testing.T) {
	testCases := []struct {
		Existed     []int
		ExpectedVal int
		ExpectedErr error
	}{
		{[]int{}, 0, errors.New("slice it empty")},
		{[]int{5, 12, 33}, 5, nil},
		{[]int{0, 5, 33}, 0, nil},
		{[]int{-2, 0, 5, 33}, -2, nil},
	}

	for _, testCase := range testCases {
		st := Slice{testCase.Existed}

		if v, err := st.GetMin(); v != testCase.ExpectedVal && err != testCase.ExpectedErr {
			t.Fatalf("Function GetMin return not min value")
		}
	}
}

func TestGetMax(t *testing.T) {
	testCases := []struct {
		Existed     []int
		ExpectedVal int
		ExpectedErr error
	}{
		{[]int{}, 0, errors.New("slice it empty")},
		{[]int{5, 12, 33}, 33, nil},
		{[]int{-33, 5, 33}, 33, nil},
	}

	for _, testCase := range testCases {
		st := Slice{testCase.Existed}

		if v, err := st.GetMax(); v != testCase.ExpectedVal && err != testCase.ExpectedErr {
			t.Fatalf("Function GetMax return not min value")
		}
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
	items := getTestSet(originalSize)
	tmpSlice := make([]int, len(items))

	for i := 0; i < b.N; i++ {
		copy(tmpSlice, items)
		ss := Slice{tmpSlice}

		ss.Insert(value)

		if originalSize >= len(ss.GetItems()) {
			b.Fatal()
		}
	}
}

func benchmarkInsertBatch(b *testing.B, size int) {
	items := getTestSet(size / 2)
	tmpSlice := make([]int, len(items))

	for i := 0; i < b.N; i++ {
		copy(tmpSlice, items)
		ss := Slice{tmpSlice}

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
		ss := Slice{tmpSlice}

		ss.Delete(value)

		if len(items) <= len(ss.GetItems()) {
			b.Fail()
		}
	}
}

func benchmarkGetMin(b *testing.B, size int) {
	st := Slice{getTestSet(size)}

	for i := 0; i < b.N; i++ {
		if v, _ := st.GetMin(); v != 0 {
			b.Fail()
		}
	}
}

func benchmarkGetMax(b *testing.B, size int) {
	st := Slice{getTestSet(size)}

	for i := 0; i < b.N; i++ {
		if v, _ := st.GetMax(); v != size-1 {
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
