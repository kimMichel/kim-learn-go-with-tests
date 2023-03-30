package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numToSum ...[]int) (result []int) {
	for _, numbers := range numToSum {
		result = append(result, Sum(numbers))
	}
	return
}

func SumAllTheRest(numToSum ...[]int) (result []int) {
	for _, numbers := range numToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			final := numbers[1:] // Get first and all
			result = append(result, Sum(final))
		}
	}
	return
}
