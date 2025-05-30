package sum

func Sum(slice []int) int {
	// arrLen := len(slice)
	result := 0
	// for idx := 0; idx < arrLen; idx++ {
	// 	result += slice[idx]
	// }
	for _, val := range slice {
		result += val
	}
	return result
}

func SumAll(sliceToSum ...[]int) []int {
	result := []int{}
	for _, slice := range sliceToSum {
		result = append(result, Sum(slice))
	}
	return result
}

func SumAllTail(sliceToSum ...[]int) []int {
	result := make([]int, 0, len(sliceToSum))
	for _, slice := range sliceToSum {
		//result = append(result, Sum(slice))
		// head, tail
		if len(slice) == 0 {
			result = append(result, 0)
		} else {
			tail := slice[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}
