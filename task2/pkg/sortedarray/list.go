package sortedarray

import "container/list"

func NewList() *list.List {
	return list.New()
}

func UpdateList(sList *list.List, value int) {
	if value > 0 {
		InsertToList(sList, value)
	} else {
		DeleteFromList(sList, -value)
	}
}

func InsertToList(sList *list.List, value int) {
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

func DeleteFromList(sList *list.List, value int) {
	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok {
			if value == v {
				sList.Remove(item)
				return
			}
		}
	}
}

func ListAsSlice(sList *list.List) []int {
	res := make([]int, 0, sList.Len())

	for item := sList.Front(); item != nil; item = item.Next() {
		if v, ok := item.Value.(int); ok {
			res = append(res, v)
		}
	}

	return res
}
