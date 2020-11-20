package main

import (
	"fmt"
	"task2/pkg/sorted/list"
	"task2/pkg/sorted/slice"
)

func main() {
	var inputValue int
	sSlice := slice.New()
	sList := list.New()

	for {
		if _, err := fmt.Scan(&inputValue); err != nil {
			break
		}

		sSlice.Update(inputValue)
		list.Update(sList, inputValue)

		fmt.Printf("\u001B[H\u001B[2J\nSlice: %v\nList: %v\n", sSlice.GetItems(), list.ToSlice(sList))
	}
}
