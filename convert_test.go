package rgbtodmc

import (
	"testing"
)

func TestFindNearest_SalmonVeryLight(t *testing.T) {
	InitColors()
	if nearest := FindNearest(255, 226, 226); nearest.Dmc != "3713" {
		t.Fail()
	}
}

func TestFindNearest_SnowWhite(t *testing.T) {
	InitColors()
	if nearest := FindNearest(255, 255, 255); nearest.Dmc != "B5200" {
		t.Fail()
	}
}

func TestFindNearest_CyclamenPinkDark(t *testing.T) {
	InitColors()
	if nearest := FindNearest(255, 0, 255); nearest.Dmc != "3804" {
		t.Fail()
	}
}

func TestFindNearest_Black(t *testing.T) {
	InitColors()
	if nearest := FindNearest(0, 0, 0); nearest.Dmc != "310" {
		t.Fail()
	}
}
