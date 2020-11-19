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

		// Clear the output and print current sSlice
		fmt.Printf(
			"\u001B[H\u001B[2J\nSlice: %v\nStruct: %v\nList: %v\n",
			sSlice,
			sStructure.GetItems(),
			list.ToSlice(sList),
		)
	}
}
