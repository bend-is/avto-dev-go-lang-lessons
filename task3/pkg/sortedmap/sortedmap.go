package sortedmap

import (
	"sort"
	"sync"
)

type ItemOrder struct {
	line, position int
}

func (i *ItemOrder) IsLess(j *ItemOrder) bool {
	return i.line < j.line || (i.line == j.line && i.position < j.position)
}

type SortedMap struct {
	sync.Mutex
	itemCounter map[string]int
	itemOrder   map[string]*ItemOrder
	stopWords   map[string]struct{}
}

func New() *SortedMap {
	return &SortedMap{
		itemCounter: make(map[string]int),
		itemOrder:   make(map[string]*ItemOrder),
		stopWords:   make(map[string]struct{}),
	}
}

func NewItemOrder(line, position int) *ItemOrder {
	return &ItemOrder{line: line, position: position}
}

func (s *SortedMap) AddStopWord(word string) {
	s.Lock()
	defer s.Unlock()
	s.stopWords[word] = struct{}{}
}

func (s *SortedMap) IsStopWord(word string) bool {
	s.Lock()
	defer s.Unlock()

	_, exist := s.stopWords[word]

	return exist
}

func (s *SortedMap) Delete(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.itemCounter, item)
	delete(s.itemOrder, item)
}

func (s *SortedMap) IncrementCount(item string, itemOrder *ItemOrder) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.itemCounter[item]; ok {
		s.itemCounter[item]++
		if itemOrder.IsLess(s.itemOrder[item]) {
			s.itemOrder[item] = itemOrder
		}
		return
	}
	s.addItem(item, itemOrder)
}

func (s *SortedMap) GetCount(item string) int {
	s.Lock()
	defer s.Unlock()
	return s.itemCounter[item]
}

func (s *SortedMap) GetTop(count int) []string {
	s.Lock()
	defer s.Unlock()
	top := make([]string, count)

	if len(s.itemCounter) == 0 {
		return top
	}

	keys := make([]string, 0, len(s.itemCounter))
	for key := range s.itemCounter {
		if _, exist := s.stopWords[key]; exist {
			continue
		}
		keys = append(keys, key)
	}

	// Sort keys by appearance in text.
	sort.Slice(keys, func(i, j int) bool {
		return s.itemOrder[keys[i]].IsLess(s.itemOrder[keys[j]])
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

func (s *SortedMap) addItem(item string, itemOrder *ItemOrder) {
	if _, ok := s.itemCounter[item]; !ok {
		s.itemCounter[item] = 1
		s.itemOrder[item] = itemOrder
	}
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
