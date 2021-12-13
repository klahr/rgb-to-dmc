package rgbtodmc

import (
	"fmt"
	"math"
)

func FindNearest(red, green, blue uint8) Color {
	if len(Colors) == 0 {
		panic("Init() has not been called")
	}

	var bestMatch Color
	bestD := 1000.0
	for _, color := range Colors {
		r, g, b := (float64)(red-color.Red), (float64)(green-color.Green), (float64)(blue-color.Blue)
		var d float64

		if ((int32)(red)+(int32)(color.Red))/2 < 128 {
			d = math.Sqrt(2.0*r*r + 4.0*g*g + 3.0*b*b)
		} else {
			d = math.Sqrt(3.0*r*r + 4.0*g*g + 2.0*b*b)
		}
		fmt.Printf("%f %s\n", d, color.Name)
		if d < bestD {
			bestD = d
			bestMatch = color
		}
	}

	return bestMatch
}
