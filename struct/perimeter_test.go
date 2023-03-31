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

	testsArea := []struct {
		form   Form
		expect float64
	}{
		{form: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{form: Circle{Radius: 10}, expect: 314.1592653589793},
		{form: Triangle{Height: 12, Base: 6}, expect: 36.0},
	}

	for _, tt := range testsArea {
		result := tt.form.Area()
		if result != tt.expect {
			t.Errorf("%#v resultado %.2f, esperado %.2f", tt.form, result, tt.expect)
		}
	}

}
