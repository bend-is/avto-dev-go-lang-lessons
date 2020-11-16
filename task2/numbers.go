package main

import "fmt"

func main() {
	var numbers []int
	var inputNumber int

	for {
		if _, err := fmt.Scan(&inputNumber); err != nil {
			break
		}

		numbers = UpdateNumbers(numbers, inputNumber)

		// Clear the output and print current numbers
		fmt.Printf("\u001B[H\u001B[2J\n%v\n", numbers)
	}
}

//UpdateNumbers update slice of numbers. If new number is positive it will be inserted else deleted from slice.
func UpdateNumbers(numbers []int, number int) []int {
	if number > 0 {
		numbers = Insert(numbers, number)
	} else {
		numbers = Delete(numbers, -number)
	}

	return numbers
}

//Insert value into slice with sorted order.
func Insert(sl []int, value int) []int {
	for i, v := range sl {
		if value < v {
			return append(sl[:i], append([]int{value}, sl[i:]...)...)
		}
	}

	return append(sl, value)
}

//InsertV2 with preallocate.
func InsertV2(sl []int, value int) []int {
	newSlice := make([]int, len(sl)+1)
	copy(newSlice, sl)

	for i, v := range sl {
		if value < v {
			copy(newSlice[i+1:], newSlice[i:])
			newSlice[i] = value

			return newSlice
		}
	}

	newSlice[len(sl)] = value

	return newSlice
}

//InsertV3 with double size allocate.
func InsertV3(sl []int, value int) []int {
	for i, v := range sl {
		if value < v {
			sl = append(sl, 0)
			copy(sl[i+1:], sl[i:])
			sl[i] = value

			return sl
		}
	}

	return append(sl, value)
}

//Delete value from slice.
func Delete(sl []int, value int) []int {
	for i, v := range sl {
		if v == value {
			return append(sl[:i], sl[i+1:]...)
		}
	}

	return sl
}
