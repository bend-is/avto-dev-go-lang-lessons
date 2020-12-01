package sortedmap

import (
	"sort"
	"sync"
)

type SortedMap struct {
	sync.Mutex
	itemCounter map[string]int
	itemOrder   map[string]int
}

func New() *SortedMap {
	return &SortedMap{
		itemCounter: make(map[string]int),
		itemOrder:   make(map[string]int),
	}
}

func (s *SortedMap) Add(item string, orderVal int) {
	if _, ok := s.itemCounter[item]; !ok {
		s.itemCounter[item] = 1
		s.itemOrder[item] = orderVal
	}
}

func (s *SortedMap) Delete(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.itemCounter, item)
	delete(s.itemOrder, item)
}

func (s *SortedMap) IncrementCount(item string, orderVal int) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.itemCounter[item]; ok {
		s.itemCounter[item]++
		if orderVal < s.itemOrder[item] {
			s.itemOrder[item] = orderVal
		}
		return
	}
	s.Add(item, orderVal)
}

func (s *SortedMap) GetCount(item string) int {
	return s.itemCounter[item]
}

func (s *SortedMap) GetTop(count int) []string {
	top := make([]string, count)

	if len(s.itemCounter) == 0 {
		return top
	}

	keys := make([]string, 0, len(s.itemCounter))
	for key := range s.itemCounter {
		keys = append(keys, key)
	}

	// Sort keys by appearance in text.
	sort.Slice(keys, func(i, j int) bool {
		return s.itemOrder[keys[i]] < s.itemOrder[keys[j]]
	})

	if len(keys) <= count {
		copy(top, keys)
		return top
	}

	copy(top, keys[:count])
	topMinI, topMinV := s.getValueWithMinCount(top)

	for _, item := range keys[count:] {
		if s.itemCounter[item] > s.itemCounter[topMinV] {
			if topMinI < len(top)-1 {
				copy(top[topMinI:], top[topMinI+1:])
			}
			top[len(top)-1] = item

			// Recalculate new topMinV and topMinI
			topMinI, topMinV = s.getValueWithMinCount(top)
		}
	}

	return top
}

// nolint
func (s *SortedMap) getValueWithMinCount(top []string) (int, string) {
	minI, minV := 0, top[0]
	for i, v := range top {
		if s.itemCounter[v] <= s.itemCounter[minV] {
			minI, minV = i, v
		}
	}

	return minI, minV
}
