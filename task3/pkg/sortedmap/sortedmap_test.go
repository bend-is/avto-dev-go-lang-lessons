package sortedmap

import (
	"math/rand"
	"strconv"
	"testing"
)

type TestCase struct {
	obj      *SortedMap
	expected []string
}

func TestSortedMap_GetTop10(t *testing.T) {
	testCases := []TestCase{
		{
			obj: &SortedMap{
				itemOrder:   map[string]*ItemOrder{},
				itemCounter: map[string]int{},
			},
			expected: []string{"", "", "", "", "", "", "", "", "", ""},
		},
		{
			obj: &SortedMap{
				itemOrder: map[string]*ItemOrder{
					"word1": {1, 1},
					"word2": {1, 2},
					"word3": {1, 3},
					"word4": {1, 4},
					"word5": {1, 5},
					"word6": {1, 6},
					"word7": {1, 7},
				},
				itemCounter: map[string]int{"word1": 1, "word2": 2, "word3": 14, "word4": 2, "word5": 12, "word6": 2, "word7": 1},
			},
			expected: []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7", "", "", ""},
		},
		{
			obj: &SortedMap{
				itemOrder: map[string]*ItemOrder{
					"word1":  {1, 1},
					"word2":  {1, 2},
					"word3":  {1, 3},
					"word4":  {1, 4},
					"word5":  {2, 5},
					"word6":  {2, 6},
					"word7":  {2, 7},
					"word8":  {3, 8},
					"word9":  {3, 9},
					"word10": {3, 10},
					"word11": {3, 11},
					"word12": {5, 12},
					"word13": {5, 13},
					"word14": {6, 14},
					"word15": {7, 15},
					"word16": {8, 16},
				},
				itemCounter: map[string]int{
					"word1":  1,
					"word2":  2,
					"word3":  14,
					"word4":  2,
					"word5":  12,
					"word6":  2,
					"word7":  1,
					"word8":  2,
					"word9":  150,
					"word10": 2,
					"word11": 2,
					"word12": 12,
					"word13": 13,
					"word14": 1,
					"word15": 6,
					"word16": 1,
				},
			},
			expected: []string{"word2", "word3", "word4", "word5", "word6", "word8", "word9", "word12", "word13", "word15"},
		},
	}

	for _, tt := range testCases {
		top := tt.obj.GetTop(10)
		if cap(top) > 10 {
			t.Fatalf("top 10 method exceeded capacity")
		}
		for i, v := range top {
			if tt.expected[i] != v {
				t.Fatalf("unexpected word at index %d: wont %s - got %s", i, tt.expected[i], v)
			}
		}
	}
}

func BenchmarkSortedMap_GetTop10(b *testing.B) {
	itemCounter := make(map[string]int)
	itemOrder := make(map[string]*ItemOrder)
	rand.Seed(1)

	for i := 0; i < 10000; i++ {
		word := "word" + strconv.Itoa(i)
		itemCounter[word] = rand.Intn(100) //nolint
		itemOrder[word] = &ItemOrder{i, i}
	}

	sMap := &SortedMap{itemOrder: itemOrder, itemCounter: itemCounter}

	for i := 0; i < b.N; i++ {
		_ = sMap.GetTop(10)
	}
}
