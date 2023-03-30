package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	verifyErrorMessage := func(t *testing.T, result, expect int, numbers []int) {
		t.Helper()
		if result != expect {
			t.Errorf("result '%d', expected '%d', passed '%v'", result, expect, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expect := 15

		verifyErrorMessage(t, result, expect, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expect := 6

		verifyErrorMessage(t, result, expect, numbers)
	})

}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{1, 2, 3})
	expect := []int{3, 6}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result '%d', expect'%d'", result, expect)
	}
}

func TestSumAllTheRest(t *testing.T) {

	verifySum := func(t *testing.T, result, expect []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("result '%v', expect '%v'", result, expect)
		}
	}
	t.Run("do sum of slices with values", func(t *testing.T) {
		result := SumAllTheRest([]int{1, 2}, []int{0, 9})
		expect := []int{2, 9}

		verifySum(t, result, expect)
	})

	t.Run("sum with empty slices", func(t *testing.T) {
		result := SumAllTheRest([]int{}, []int{3, 4, 5})
		expect := []int{0, 9}

		verifySum(t, result, expect)
	})

}
