package arrayAndSlice

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSumTails ...[]int) (tailsSums []int) {
	for _, numbers := range numbersToSumTails {
		if len(numbers) > 0 {
			numbers = numbers[1:]
		}
		tailsSums = append(tailsSums, Sum(numbers))
	}
	return
}