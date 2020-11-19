package list

import "container/list"

func New() *list.List {
	return list.New()
}

func Update(sList *list.List, value int) {
	if value > 0 {
		Insert(sList, value)
	} else {
		Delete(sList, -value)
	}
}

func Insert(sList *list.List, value int) {
	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok {
			if value < v {
				sList.InsertBefore(value, item)
				return
			}
		}
	}

	sList.PushBack(value)
}

func Delete(sList *list.List, value int) {
	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok {
			if value == v {
				sList.Remove(item)
				return
			}
		}
	}
}

func ToSlice(sList *list.List) []int {
	res := make([]int, 0, sList.Len())

	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok {
			res = append(res, v)
		}
	}

	return res
}

func GetMax(sList *list.List) int {
	var max int

	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok && (max == 0 || v > max) {
			max = v
		}
	}

	return max
}

func GetMin(sList *list.List) int {
	var min int

	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok && (min == 0 || v < min) {
			min = v
		}
	}

	return min
}
