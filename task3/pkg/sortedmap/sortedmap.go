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
	s.items = append(s.items, item)
	s.itemCounter[item] = 1
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

func (s *SortedMap) GetTop10() []string {
	count := 10
	top := make([]string, count)

	if len(s.items) < count {
		copy(top, s.items)
		return top
	}
	copy(top, s.items[:count])

	for _, item := range s.items[count:] {
		for topI, topV := range top {
			if s.itemCounter[topV] < s.itemCounter[item] {
				if topI != count-1 {
					copy(top[topI:], top[topI+1:])
				}
				top[count-1] = item
				break
			}
		}
	}

	return top
}