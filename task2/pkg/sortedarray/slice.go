package sortedarray

//Update slice. If new number is positive it will be inserted else deleted from slice.
func UpdateSlice(sl []int, value int) []int {
	if value > 0 {
		sl = InsertToSliceV3(sl, value)
	} else {
		sl = DeleteFromSlice(sl, -value)
	}

	return sl
}

//InsertToSlice value into slice with sorted order.
func InsertToSlice(sl []int, value int) []int {
	for i, v := range sl {
		if value < v {
			return append(sl[:i], append([]int{value}, sl[i:]...)...)
		}
	}

	return append(sl, value)
}

//InsertToSliceV2 with preallocate.
func InsertToSliceV2(sl []int, value int) []int {
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

//InsertToSliceV3 with double size allocate.
func InsertToSliceV3(sl []int, value int) []int {
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
func DeleteFromSlice(sl []int, value int) []int {
	for i, v := range sl {
		if v == value {
			return append(sl[:i], sl[i+1:]...)
		}
	}

	return sl
}

func GetSliceMax(sl []int) int {
	max := sl[len(sl)-1]

	for _, v := range sl {
		if v > max {
			max = v
		}
	}

	return max
}

func GetSliceMin(sl []int) int {
	min := sl[0]

	for _, v := range sl {
		if v < min {
			min = v
		}
	}

	return min
}
