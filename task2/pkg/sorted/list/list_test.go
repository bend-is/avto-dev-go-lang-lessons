package list

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		Values1  []int
		Values2  []int
		Expected bool
	}{
		{[]int{}, []int{}, true},
		{[]int{}, []int{1}, false},
		{[]int{1}, []int{1}, true},
		{[]int{1}, []int{2}, false},
		{[]int{1, 4, 10}, []int{1, 4, 10}, true},
		{[]int{1, 5, 10}, []int{1, 4, 10}, false},
	}

	for _, v := range testCases {
		lst1 := New()
		lst2 := New()
		lst1.Update(v.Values1...)
		lst2.Update(v.Values2...)

		if lst1.Equal(*lst2) != v.Expected {
			t.Fatalf("Got wrong result while comparing two objects: %v, %v", lst1.GetItems(), lst2.GetItems())
		}
	}
}

func TestInsert(t *testing.T) {
	lst := New()

	testCases := []struct {
		Value    int
		Expected []int
	}{
		{0, []int{0}},
		{1, []int{0, 1}},
		{3, []int{0, 1, 3}},
		{2, []int{0, 1, 2, 3}},
		{2, []int{0, 1, 2, 2, 3}},
		{-3, []int{-3, 0, 1, 2, 2, 3}},
	}

	for _, v := range testCases {
		lst.Insert(v.Value)

		if !isSliceEquals(lst.GetItems(), v.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", v.Expected, lst.GetItems())
		}
	}
}

func TestDelete(t *testing.T) {
	lst := New()

	if lst.Delete(5); !isSliceEquals(lst.GetItems(), []int{}) {
		t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", []int{}, lst.GetItems())
	}

	// Fill list with data.
	for _, v := range []int{-3, 1, 4, 4, 5, 6} {
		lst.Insert(v)
	}

	testCases := []struct {
		Value    int
		Expected []int
	}{
		{1, []int{-3, 4, 4, 5, 6}},
		{99, []int{-3, 4, 4, 5, 6}},
		{4, []int{-3, 4, 5, 6}},
		{-3, []int{4, 5, 6}},
	}

	for _, v := range testCases {
		lst.Delete(v.Value)

		if !isSliceEquals(lst.GetItems(), v.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", v.Expected, lst.GetItems())
		}
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		Values   []int
		Expected []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 1, -1}, []int{1}},
		{[]int{-1, -1, -1}, []int{}},
		{[]int{10, 4, 3, 9, -5, -3}, []int{4, 9, 10}},
	}

	for _, v := range testCases {
		lst := New()
		lst.Update(v.Values...)

		if !isSliceEquals(lst.GetItems(), v.Expected) {
			t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", v.Expected, lst.GetItems())
		}
	}
}

func TestGetMin(t *testing.T) {
	testCases := []struct {
		Values      []int
		ExpectedVal int
		ExpectedErr error
	}{
		{[]int{}, 0, errors.New("list is empty")},
		{[]int{0}, 0, nil},
		{[]int{3, 4, 9}, 3, nil},
	}

	for _, v := range testCases {
		lst := New()
		lst.Update(v.Values...)

		if min, err := lst.GetMin(); min != v.ExpectedVal && err != v.ExpectedErr {
			t.Fatalf("Function GetMin return not min value")
		}
	}
}

func TestGetMax(t *testing.T) {
	testCases := []struct {
		Values      []int
		ExpectedVal int
		ExpectedErr error
	}{
		{[]int{}, 0, errors.New("list is empty")},
		{[]int{0}, 0, nil},
		{[]int{1}, 1, nil},
		{[]int{3, 4, 9}, 9, nil},
	}

	for _, v := range testCases {
		lst := New()
		lst.Update(v.Values...)

		if min, err := lst.GetMax(); min != v.ExpectedVal && err != v.ExpectedErr {
			t.Fatalf("Function GetMax return not min value")
		}
	}
}

func BenchmarkInsertIntoTheBeginning(b *testing.B) { benchmarkInsert(b, 10000, 1) }
func BenchmarkInsertIntoMiddle(b *testing.B)       { benchmarkInsert(b, 10000, 5000) }
func BenchmarkInsertIntoTheEnd(b *testing.B)       { benchmarkInsert(b, 10000, 10000) }
func BenchmarkInsertBatch(b *testing.B)            { benchmarkInsertBatch(b, 10000) }

func BenchmarkGetMinOn100(b *testing.B)   { benchmarkGetMin(b, 100) }
func BenchmarkGetMinOn10000(b *testing.B) { benchmarkGetMin(b, 10000) }

func BenchmarkGetMaxOn100(b *testing.B)   { benchmarkGetMax(b, 100) }
func BenchmarkGetMaxOn10000(b *testing.B) { benchmarkGetMax(b, 10000) }

func BenchmarkEqual(b *testing.B) {
	items := getTestSet(1000)
	lst1 := New()
	lst2 := New()
	lst1.Update(items...)
	lst2.Update(items...)

	benchmarkEqual(b, lst1, lst2, true)
}

func BenchmarkNotEqualByLen(b *testing.B) {
	items := getTestSet(1000)
	lst1 := New()
	lst2 := New()
	lst1.Update(items...)
	lst2.Update(items[:len(items)-1]...)

	benchmarkEqual(b, lst1, lst2, false)
}

func BenchmarkNotEqualInTheEnd(b *testing.B) {
	items := getTestSet(1000)
	lst1 := New()
	lst2 := New()
	lst1.Update(items...)
	lst2.Update(items...)
	lst1.Insert(10000)
	lst2.Insert(10001)

	benchmarkEqual(b, lst1, lst2, false)
}

func BenchmarkNotEqualInTheBegin(b *testing.B) {
	items := getTestSet(1000)
	items[0] = items[1]
	lst1 := New()
	lst2 := New()
	lst1.Update(getTestSet(1000)...)
	lst2.Update(items...)

	benchmarkEqual(b, lst1, lst2, false)
}

func BenchmarkNotEqualInMiddle(b *testing.B) {
	items := getTestSet(1000)
	lst1 := New()
	lst2 := New()
	lst1.Update(items...)
	lst2.Update(items...)
	lst2.Delete(501)
	lst2.Insert(500)

	benchmarkEqual(b, lst1, lst2, false)
}

func benchmarkInsert(b *testing.B, size, value int) {
	lst := New()
	lst.Update(getTestSet(size)...)

	for i := 0; i < b.N; i++ {
		lst.Insert(value)

		if size >= lst.Len() {
			b.Fatal()
		}
	}
}

func benchmarkInsertBatch(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		lst := New()
		lst.Update(getTestSet(size / 2)...)
		b.StartTimer()

		for j := 0; j < size; j++ {
			lst.Insert(j)
		}
	}
}

func benchmarkGetMin(b *testing.B, size int) {
	lst := New()

	for _, v := range getTestSet(size) {
		lst.Insert(v)
	}

	for i := 0; i < b.N; i++ {
		min, _ := lst.GetMin()
		if min != 1 {
			b.Fail()
		}
	}
}

func benchmarkGetMax(b *testing.B, size int) {
	lst := New()

	for _, v := range getTestSet(size) {
		lst.Insert(v)
	}

	for i := 0; i < b.N; i++ {
		max, _ := lst.GetMax()
		if max != size {
			b.Fail()
		}
	}
}

func benchmarkEqual(b *testing.B, lst1, lst2 *SList, res bool) {
	for i := 0; i < b.N; i++ {
		if lst1.Equal(*lst2) != res {
			b.Fail()
		}
	}
}

func isSliceEquals(lst1, lst2 []int) bool {
	if len(lst1) != len(lst2) {
		return false
	}

	for i, v := range lst1 {
		if v != lst2[i] {
			return false
		}
	}

	return true
}

func getTestSet(size int) []int {
	testSet := make([]int, 0, size+1)

	for i := 1; i < size+1; i++ {
		testSet = append(testSet, i)
	}

	return testSet
}
