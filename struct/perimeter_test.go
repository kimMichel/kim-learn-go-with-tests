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
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range testsArea {
		result := tt.form.Area()
		if result != tt.expect {
			t.Errorf("resultado %.2f, esperado %.2f", result, tt.expect)
		}
	}

}
