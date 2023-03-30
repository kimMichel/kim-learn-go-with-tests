package main

import (
	"reflect"
	"testing"
)

func TestRecapSum(t *testing.T) {

	verifyNumError := func(t *testing.T, result, expect int) {
		t.Helper()
		if result != expect {
			t.Errorf("result e'%d', expect '%d'", result, expect)
		}
	}

	t.Run("soma dos valores dentro do slice", func(t *testing.T) {
		result := RecapSum([]int{1, 2, 3, 4, 5})
		expect := 15

		verifyNumError(t, result, expect)
	})

	t.Run("soma dos valores caso o slice esteja vazio", func(t *testing.T) {
		result := RecapSum([]int{})
		expect := 0

		verifyNumError(t, result, expect)
	})
}

func TestRecapSumAllValues(t *testing.T) {

	verifySliceError := func(t *testing.T, result, expect []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("result '%v', expect '%v'", result, expect)
		}
	}

	t.Run("soma de todos os valores", func(t *testing.T) {
		result := RecapSumAllValues([]int{1, 2}, []int{0, 9})
		expect := []int{3, 9}

		verifySliceError(t, result, expect)
	})

	t.Run("soma de todos os valores caso tenha vazio", func(t *testing.T) {
		result := RecapSumAllValues([]int{0, 0}, []int{}, []int{1, 2, 3, 4, 5})
		expect := []int{0, 0, 15}

		verifySliceError(t, result, expect)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecapSumAllValues([]int{0, 0}, []int{}, []int{1, 2, 3, 4, 5})
	}
}

func TestRecapSumAllValuesAfterFirst(t *testing.T) {

	verifyError := func(t *testing.T, result, expect []int) {
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("result '%v', expect '%v'", result, expect)
		}
	}

	t.Run("soma quando todos os valores estiverem preenchidos", func(t *testing.T) {
		result := RecapSumAllValuesAfterFirst([]int{1, 2, 3}, []int{0, 9})
		expect := []int{5, 9}

		verifyError(t, result, expect)
	})

	t.Run("soma quando os valores forem vazios", func(t *testing.T) {
		result := RecapSumAllValuesAfterFirst([]int{}, []int{0, 9})
		expect := []int{0, 9}

		verifyError(t, result, expect)
	})

}

func TestSumJustTwoFirst(t *testing.T) {
	result := SumJustTwoFirst([]int{1, 2, 3}, []int{2, 3}, []int{3, 3, 3})
	expect := []int{3, 5, 6}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result '%v', expect '%v'", result, expect)
	}
}
