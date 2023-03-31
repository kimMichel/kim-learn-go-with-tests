package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expect := 40.0

	if result != expect {
		t.Errorf("result '%.2f', expect '%.2f'", result, expect)
	}
}

func TestArea(t *testing.T) {

	verifyFloatError := func(t *testing.T, form Form, expect float64) {
		t.Helper()
		result := form.Area()
		if result != expect {
			t.Errorf("result '%.2f', expect '%.2f'", result, expect)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expect := 72.0

		verifyFloatError(t, rectangle, expect)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		expect := 314.1592653589793

		verifyFloatError(t, circle, expect)
	})
}
