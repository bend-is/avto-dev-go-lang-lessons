package slice

import "errors"

type Slice struct {
	items []int
}

func New() *Slice {
	return new(Slice)
}

//Update structure. If new item is positive it will be inserted else deleted from slice.
func (s *Slice) Update(values ...int) {
	for _, value := range values {
		if value > 0 {
			s.Insert(value)
		} else {
			s.Delete(-value)
		}
	}
}

//Insert value to structure
func (s *Slice) Insert(value int) {
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
func (s *Slice) Delete(value int) {
	for i, v := range s.items {
		if v == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
			break
		}
	}
}

func (s *Slice) GetItems() []int {
	// Prevent slice from changes outside.
	tmp := make([]int, len(s.items))
	copy(tmp, s.items)

	return tmp
}

func (s *Slice) GetMax() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("slice is empty")
	}

	// Cause access to items is private and we knows that is defiantly sorted.
	return s.items[len(s.items)-1], nil
}

func (s *Slice) GetMin() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("slice is empty")
	}

	// Cause access to items is private and we knows that is defiantly sorted.
	return s.items[0], nil
}
