package sortedmap

type SortedMap struct {
	items       []string
	itemCounter map[string]int
}

func New() *SortedMap {
	return &SortedMap{
		itemCounter: make(map[string]int),
	}
}

func (s *SortedMap) Add(item string) {
	if _, ok := s.itemCounter[item]; !ok {
		s.items = append(s.items, item)
		s.itemCounter[item] = 1
	}
}

func (s *SortedMap) Delete(item string) {
	if itemsLen := len(s.items); itemsLen > 0 && s.items[itemsLen-1] == item {
		s.items = s.items[:itemsLen-1]
		delete(s.itemCounter, item)
		return
	}

	for i, v := range s.items {
		if v == item {
			s.items = append(s.items[:i], s.items[i+1:]...)
			delete(s.itemCounter, item)
			break
		}
	}
}

func (s *SortedMap) IncrementCount(item string) {
	if _, ok := s.itemCounter[item]; ok {
		s.itemCounter[item]++
		return
	}

	s.Add(item)
}

func (s *SortedMap) GetCount(item string) int {
	return s.itemCounter[item]
}

func (s *SortedMap) GetTop(count int) []string {
	top := make([]string, count)

	if len(s.items) < count {
		copy(top, s.items)
		return top
	}
	copy(top, s.items[:count])

	topMinI, topMinV := s.getValueWithMinCount(top)

	for _, item := range s.items[count:] {
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

func (s *SortedMap) getValueWithMinCount(top []string) (int, string) {
	minI, minV := 0, top[0]
	for i, v := range top {
		if s.itemCounter[v] < s.itemCounter[minV] {
			minI, minV = i, v
		}
	}

	return minI, minV
}
