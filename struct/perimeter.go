package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() (result float64) {
	return (r.Height * r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (result float64) {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() (result float64) {
	return (t.Base * t.Height) * 0.5
}

type Form interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) (result float64) {
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) (result float64) {
	return (rectangle.Width * rectangle.Height)
}
