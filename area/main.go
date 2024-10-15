package area

import (
	"math"
)

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Area() float64 {
	return math.Pow(r.Radius, 2) * math.Pi
}
