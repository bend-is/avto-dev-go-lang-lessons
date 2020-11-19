package sortedarray

type SortedStruct struct {
	items []int
}

func NewStruct() *SortedStruct {
	return new(SortedStruct)
}

//Update struct. If new item is positive it will be inserted else deleted from slice.
func (s *SortedStruct) Update(value int) {
	if value > 0 {
		s.Insert(value)
	} else {
		s.Delete(-value)
	}
}

//Insert value to struct
func (s *SortedStruct) Insert(value int) {
	for i, v := range s.items {
		if value < v {
			s.items = append(s.items, 0)
			copy(s.items[i+1:], s.items[i:])
			s.items[i] = value

			return
		}
	}

	s.items = append(s.items, value)
}

//Delete value from struct.
func (s *SortedStruct) Delete(value int) {
	for i, v := range s.items {
		if v == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
		}
	}
}

func (s *SortedStruct) GetItems() []int {
	tmp := make([]int, len(s.items))
	copy(tmp, s.items)

	return tmp
}

func (s *SortedStruct) GetMax() int {
	return s.items[len(s.items)-1]
}

func (s *SortedStruct) GetMin() int {
	return s.items[0]
}
