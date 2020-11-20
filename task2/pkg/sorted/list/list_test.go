package list

import "testing"

func TestInsert(t *testing.T) {
	lst := New()
	expected := []int{1, 4, 5, 6}

	for _, v := range expected {
		Insert(lst, v)
	}

	if !isSliceEquals(ToSlice(lst), expected) {
		t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", expected, ToSlice(lst))
	}
}

func TestDelete(t *testing.T) {
	lst := New()
	expected := []int{4, 6}

	for _, v := range []int{1, 4, 5, 6} {
		Insert(lst, v)
	}

	Delete(lst, 1)
	Delete(lst, 5)

	if !isSliceEquals(ToSlice(lst), expected) {
		t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", expected, ToSlice(lst))
	}
}

func TestUpdate(t *testing.T) {
	lst := New()
	input := []int{10, 4, 3, 9, -5, -3}
	expected := []int{4, 9, 10}

	for _, v := range input {
		Update(lst, v)
	}

	if !isSliceEquals(ToSlice(lst), expected) {
		t.Fatalf("Failed compare that two slices are equals. Want %v, got %v", expected, ToSlice(lst))
	}
}

func TestGetMin(t *testing.T) {
	lst := New()

	for _, v := range []int{10, 4, 3, 9} {
		Insert(lst, v)
	}

	if GetMin(lst) != 3 {
		t.Fatalf("Function GetMin return not min value")
	}
}

func TestGetMax(t *testing.T) {
	lst := New()

	for _, v := range []int{10, 4, 3, 9} {
		Insert(lst, v)
	}

	if GetMax(lst) != 10 {
		t.Fatalf("Function GetMax return not max value")
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

func benchmarkInsert(b *testing.B, size, value int) {
	lst := New()
	for _, v := range getTestSet(size) {
		Insert(lst, v)
	}

	for i := 0; i < b.N; i++ {
		Insert(lst, value)

		if size >= lst.Len() {
			b.Fatal()
		}
	}
}

func benchmarkInsertBatch(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		lst := New()
		for _, v := range getTestSet(size / 2) {
			Insert(lst, v)
		}
		b.StartTimer()

		for j := 0; j < size; j++ {
			Insert(lst, j)
		}
	}
}

func benchmarkGetMin(b *testing.B, size int) {
	lst := New()

	for _, v := range getTestSet(size) {
		Insert(lst, v)
	}

	for i := 0; i < b.N; i++ {
		if GetMin(lst) != 1 {
			b.Fail()
		}
	}
}

func benchmarkGetMax(b *testing.B, size int) {
	lst := New()

	for _, v := range getTestSet(size) {
		Insert(lst, v)
	}

	for i := 0; i < b.N; i++ {
		if GetMax(lst) != size {
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
	testSet := make([]int, 0, size+1)

	for i := 1; i < size+1; i++ {
		testSet = append(testSet, i)
	}

	return testSet
}
