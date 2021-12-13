package rgbtodmc

import (
	"math"
)

func FindNearest(red, green, blue uint8) Color {
	if len(Colors) == 0 {
		panic("InitColors() has not been called")
	}

	var bestMatch Color
	bestD := 1000.0
	for _, color := range Colors {
		r, g, b := math.Abs((float64)(red)-(float64)(color.Red)), math.Abs((float64)(green)-(float64)(color.Green)), math.Abs((float64)(blue)-(float64)(color.Blue))
		var d float64

		if ((int32)(red)+(int32)(color.Red))/2 < 128 {
			d = math.Sqrt(2.0*r*r + 4.0*g*g + 3.0*b*b)
		} else {
			d = math.Sqrt(3.0*r*r + 4.0*g*g + 2.0*b*b)
		}
		if d < bestD {
			bestD = d
			bestMatch = color
		}
	}

	return bestMatch
}
