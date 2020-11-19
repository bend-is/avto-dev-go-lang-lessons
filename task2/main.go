package main

import (
	"fmt"
	"task2/pkg/sortedarray"
)

func main() {
	var inputValue int
	var sSlice []int
	sStruct := sortedarray.NewStruct()
	sList := sortedarray.NewList()

	for {
		if _, err := fmt.Scan(&inputValue); err != nil {
			break
		}

		sStruct.Update(inputValue)
		sortedarray.UpdateList(sList, inputValue)
		sSlice = sortedarray.UpdateSlice(sSlice, inputValue)

		// Clear the output and print current sSlice
		fmt.Printf(
			"\u001B[H\u001B[2J\nSlice: %v\nStruct: %v\nList: %v\n",
			sSlice,
			sStruct.GetItems(),
			sortedarray.ListAsSlice(sList),
		)
	}
}
