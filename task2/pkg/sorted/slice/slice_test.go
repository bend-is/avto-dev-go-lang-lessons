package slice

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		Slice1   Slice
		Slice2   Slice
		Expected bool
	}{
		{Slice{}, Slice{}, true},
		{Slice{}, Slice{[]int{0}}, false},
		{Slice{[]int{1}}, Slice{[]int{1}}, true},
		{Slice{[]int{1}}, Slice{[]int{2}}, false},
		{Slice{[]int{1, 4, 10}}, Slice{[]int{1, 4, 10}}, true},
		{Slice{[]int{1, 5, 10}}, Slice{[]int{1, 4, 10}}, false},
	}

	for _, v := range testCases {
		if v.Slice1.Equal(v.Slice2) != v.Expected {
			t.Fatalf("Got wrong result while comparing two objects: %v, %v", v.Slice1, v.Slice2)
		}
	}
}

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

		if !ss.Equal(Slice{testCase.Expected}) {
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

		if !ss.Equal(Slice{testCase.Expected}) {
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

		if !ss.Equal(Slice{testCase.Expected}) {
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

func TestLen(t *testing.T) {
	testCases := []struct {
		slice    Slice
		expected int
	}{
		{Slice{[]int{}}, 0},
		{Slice{[]int{1}}, 1},
		{Slice{[]int{-1, -2, 1, 4, 5}}, 5},
	}

	for _, v := range testCases {
		if v.slice.Len() != v.expected {
			t.Fatalf("Wrong response for len function. Want %v, got %v", v.expected, v.slice.Len())
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

func BenchmarkEqual(b *testing.B) {
	sl1 := Slice{getTestSet(1000)}
	sl2 := Slice{getTestSet(1000)}

	benchmarkEqual(b, sl1, sl2, true)
}

func BenchmarkNotEqualByLen(b *testing.B) {
	sl1 := Slice{getTestSet(1000)}
	sl2 := Slice{getTestSet(1000)}
	sl2.Delete(400)

	benchmarkEqual(b, sl1, sl2, false)
}

func BenchmarkNotEqualInTheEnd(b *testing.B) {
	sl1 := Slice{getTestSet(1000)}
	sl2 := Slice{getTestSet(1000)}
	sl1.Insert(10000)
	sl2.Insert(10001)

	benchmarkEqual(b, sl1, sl2, false)
}

func BenchmarkNotEqualInTheBegin(b *testing.B) {
	item1 := getTestSet(1000)
	item2 := getTestSet(1000)
	item2[0] = item2[1]
	sl1 := Slice{item1}
	sl2 := Slice{item2}

	benchmarkEqual(b, sl1, sl2, false)
}

func BenchmarkNotEqualInMiddle(b *testing.B) {
	sl1 := Slice{getTestSet(1000)}
	sl2 := Slice{getTestSet(1000)}
	sl2.Delete(501)
	sl2.Insert(500)

	benchmarkEqual(b, sl1, sl2, false)
}

func benchmarkInsert(b *testing.B, originalSize, value int) {
	items := getTestSet(originalSize)
	tmpSlice := make([]int, len(items))

	for i := 0; i < b.N; i++ {
		copy(tmpSlice, items)
		ss := Slice{tmpSlice}

		ss.Insert(value)

		if originalSize >= ss.Len() {
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

		if len(items) <= ss.Len() {
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

func benchmarkEqual(b *testing.B, sl1, sl2 Slice, res bool) {
	for i := 0; i < b.N; i++ {
		if sl1.Equal(sl2) != res {
			b.Fail()
		}
	}
}

func getTestSet(size int) []int {
	testSet := make([]int, 0, size)

	for i := 0; i < size; i++ {
		testSet = append(testSet, i)
	}

	return testSet
}
