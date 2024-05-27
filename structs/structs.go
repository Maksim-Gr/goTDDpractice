package structs

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return 0

}

func (c Circle) Area() float64 {
	return 0
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Height
}
