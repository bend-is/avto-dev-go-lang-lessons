package list

import (
	"container/list"
	"errors"
)

type SList struct {
	items list.List
}

func New() *SList {
	return new(SList)
}

func (s *SList) Front() *list.Element {
	return s.items.Front()
}

func (s *SList) Update(values ...int) {
	for _, value := range values {
		if value > 0 {
			s.Insert(value)
		} else {
			s.Delete(-value)
		}
	}
}

func (s *SList) Insert(value int) {
	if v, err := s.GetMax(); err != nil || value >= v {
		s.items.PushBack(value)
	} else if v, err := s.GetMin(); err != nil || value <= v {
		s.items.PushFront(value)
	} else {
		for item := s.items.Front(); item != nil; item = item.Next() {
			if value < item.Value.(int) {
				s.items.InsertBefore(value, item)
				break
			}
		}
	}
}

func (s *SList) Delete(value int) {
	for item := s.items.Front(); item != nil; item = item.Next() {
		if value == item.Value.(int) {
			s.items.Remove(item)
			break
		}
	}
}

func (s *SList) GetItems() []int {
	res := make([]int, 0, s.items.Len())

	for item := s.items.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok {
			res = append(res, v)
		}
	}

	return res
}

func (s *SList) GetMax() (int, error) {
	if max := s.items.Back(); max == nil {
		return 0, errors.New("list is empty")
	} else {
		return max.Value.(int), nil
	}
}

func (s *SList) GetMin() (int, error) {
	if min := s.items.Front(); min == nil {
		return 0, errors.New("list is empty")
	} else {
		return min.Value.(int), nil
	}
}

func (s *SList) Len() int {
	return s.items.Len()
}

func (s *SList) Equal(val SList) bool {
	if s.Len() != val.Len() {
		return false
	} else if s.Len() == 0 {
		return true
	}

	aMin, _ := s.GetMin()
	aMax, _ := s.GetMax()
	bMin, _ := val.GetMin()
	bMax, _ := val.GetMax()

	if aMin != bMin || aMax != bMax {
		return false
	}

	for item1, item2 := s.Front(), val.Front(); item1 != nil; item1, item2 = item1.Next(), item2.Next() {
		if item1.Value != item2.Value {
			return false
		}
	}

	return true
}
