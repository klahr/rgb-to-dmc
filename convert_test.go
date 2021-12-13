package rgbtodmc

import (
	"fmt"
	"testing"
)

func TestFindNearest(t *testing.T) {
	InitColors()
	/*if nearest := FindNearest(255, 226, 226); nearest.Dmc != "3713" {
		t.Fail()
	}
	if nearest := FindNearest(255, 255, 255); nearest.Dmc != "B5200" {
		t.Fail()
	}*/
	if nearest := FindNearest(255, 0, 255); nearest.Dmc != "" {
		fmt.Print(nearest.Dmc)
	}
}
