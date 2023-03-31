package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) (result float64) {
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) (result float64) {
	return (rectangle.Width * rectangle.Height)
}
