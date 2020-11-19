package slice

//Update slice. If new number is positive it will be inserted else deleted from slice.
func Update(sl []int, value int) []int {
	if value > 0 {
		sl = InsertV3(sl, value)
	} else {
		sl = Delete(sl, -value)
	}

	return sl
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

func GetMax(sl []int) int {
	max := sl[len(sl)-1]

	for _, v := range sl {
		if v > max {
			max = v
		}
	}

	return max
}

func GetMin(sl []int) int {
	min := sl[0]

	for _, v := range sl {
		if v < min {
			min = v
		}
	}

	return min
}
