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
				items:       []string{},
				itemCounter: map[string]int{},
			},
			expected: []string{"", "", "", "", "", "", "", "", "", ""},
		},
		{
			obj: &SortedMap{
				items: []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7"},
				itemCounter: map[string]int{
					"word1": 1,
					"word2": 2,
					"word3": 14,
					"word4": 2,
					"word5": 12,
					"word6": 2,
					"word7": 1,
				},
			},
			expected: []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7", "", "", ""},
		},
		{
			obj: &SortedMap{
				items: []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7", "word8", "word9", "word10", "word11", "word12", "word13", "word14", "word15"},
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
				},
			},
			expected: []string{"word3", "word5", "word7", "word8", "word9", "word10", "word11", "word12", "word13", "word15"},
		},
	}

	for _, tt := range testCases {
		top := tt.obj.GetTop10()
		if cap(top) > 10 {
			t.Fatalf("top 10 method exceeded capacity")
		}
		for i, v := range top {
			if tt.expected[i] != v {
				t.Fatalf("unexpected word at index %v: wont %v - got %v", i, tt.expected[i], v)
			}
		}

	}
}

func BenchmarkSortedMap_GetTop10(b *testing.B) {
	items := make([]string, 10000)
	itemCounter := make(map[string]int)
	rand.Seed(1)

	for i := 0; i < 10000; i++ {
		word := "word" + strconv.Itoa(i)
		items = append(items, word)
		itemCounter[word] = rand.Intn(100)
	}

	sMap := &SortedMap{items: items, itemCounter: itemCounter}

	for i := 0; i < b.N; i++ {
		_ = sMap.GetTop10()
	}
}
