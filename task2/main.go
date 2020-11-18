package main

import (
	"fmt"
	"task2/pkg/sortedarray"
)

func main() {
	var sortedArray []int
	var inputValue int

	for {
		if _, err := fmt.Scan(&inputValue); err != nil {
			break
		}

		sortedArray = sortedarray.Update(sortedArray, inputValue)

		// Clear the output and print current sortedArray
		fmt.Printf("\u001B[H\u001B[2J\n%v\n", sortedArray)
	}
}
