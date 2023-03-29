package main

import "testing"

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
