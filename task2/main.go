package main

import (
	"fmt"
	"task2/pkg/sorted/list"
	"task2/pkg/sorted/slice"
	"task2/pkg/sorted/structure"
)

func main() {
	var inputValue int
	var sSlice []int
	sStructure := structure.New()
	sList := list.New()

	for {
		if _, err := fmt.Scan(&inputValue); err != nil {
			break
		}

		sStructure.Update(inputValue)
		list.Update(sList, inputValue)
		sSlice = slice.Update(sSlice, inputValue)

		fmt.Printf(
			"Slice: %v\nStructure: %v\nList: %v\n\n",
			sSlice,
			sStructure.GetItems(),
			list.ToSlice(sList),
		)
		fmt.Printf(
			"Min:\nSlice: %d\nStructure: %d\nList: %d\n\n",
			slice.GetMin(sSlice),
			sStructure.GetMin(),
			list.GetMin(sList),
		)
		fmt.Printf(
			"Max:\nSlice: %d\nStructure: %d\nList: %d",
			slice.GetMax(sSlice),
			sStructure.GetMax(),
			list.GetMax(sList),
		)
	}
}
