package main

func RecapSum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func RecapSumAllValues(numToSum ...[]int) (result []int) {
	for _, numbers := range numToSum {
		// já que eles possuem uma capacidade fixa, é possível criar novos slices de antigos usando append
		result = append(result, RecapSum(numbers))
	}
	return
}

func RecapSumAllValuesAfterFirst(numToSum ...[]int) (result []int) {
	for _, numbers := range numToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			final := numbers[1:]
			result = append(result, Sum(final))
		}
	}
	return
}

func SumJustTwoFirst(numToSum ...[]int) (result []int) {
	for _, numbers := range numToSum {
		final := numbers[:2]
		result = append(result, Sum(final))
	}
	return
}
