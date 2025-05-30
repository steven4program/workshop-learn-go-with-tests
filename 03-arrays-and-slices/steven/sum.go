package main

func Sum(slice []int) int {
	sum := 0
	for _, number := range slice {
		sum += number
	}
	return sum
}

func SumAll(sliceToSum ...[]int) []int {
	var sums []int
	for _, numbers := range sliceToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(sliceToSum ...[]int) []int {
	var sums []int
	for _, numbers := range sliceToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}