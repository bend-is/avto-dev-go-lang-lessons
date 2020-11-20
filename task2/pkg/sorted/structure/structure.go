package structure

type Structure struct {
	items []int
}

func New() *Structure {
	return new(Structure)
}

//Update structure. If new item is positive it will be inserted else deleted from slice.
func (s *Structure) Update(value int) {
	if value > 0 {
		s.Insert(value)
	} else {
		s.Delete(-value)
	}
}

//Insert value to structure
func (s *Structure) Insert(value int) {
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

//Delete value from structure.
func (s *Structure) Delete(value int) {
	for i, v := range s.items {
		if v == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
		}
	}
}

func (s *Structure) GetItems() []int {
	// Prevent slice from changes outside.
	tmp := make([]int, len(s.items))
	copy(tmp, s.items)

	return tmp
}

func (s *Structure) GetMax() int {
	// Cause access to items is private and we knows that is defiantly sorted.
	return s.items[len(s.items)-1]
}

func (s *Structure) GetMin() int {
	// Cause access to items is private and we knows that is defiantly sorted.
	return s.items[0]
}
