package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numToSum ...[]int) (result []int) {
	numLength := len(numToSum)
	result = make([]int, numLength)

	for i, number := range numToSum {
		result[i] = Sum(number)
	}
	return
}
