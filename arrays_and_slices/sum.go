package arrays_and_slices

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

//func SumAll(numbersToSum ...[]int) (sums []int) {
//	for _, numbers := range numbersToSum {
//		sums = append(sums, Sum(numbers))
//		//sums[i] = Sum(numbers)
//	}
//	return
//}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
