package tinymurky

func Sum1(arr [5]int) int {
	var sum int

	for _, n := range arr {
		sum += n
	}
	return sum
}

func Sum2(arr [5]int) int {
	l := len(arr)
	var sum int
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	return sum
}

func SumSlice(slice []int) int {
	var sum int
	for _, n := range slice{
		sum += n
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	output := make([]int, 0, len(slices))

	// for i, slice := range slices {
	// 	sum := SumSlice(slice)
	// 	output[i] = sum
	// }
	for _, slice := range slices {
		sum := SumSlice(slice)
		output = append(output, sum)
	}
	return output
}

// 不加第一個
func SumAllTail(slices ...[]int) []int {
	output := make([]int, 0, len(slices))

	for _, slice := range slices {
		//sum := SumSlice(slice[1:])
		if len(slice) == 0 {
			output = append(output, 0)
			continue
		}
		
		tail := slice[1:]
		sum := SumSlice(tail)
		output = append(output, sum)
	}
	return output
}