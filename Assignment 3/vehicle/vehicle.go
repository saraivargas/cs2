package vehicle

import "strconv"

type Vehicle struct {
	VehicleType string
	Brand       string
	Color       string
	Speed       int
	Type        string
	Doors       int
	Load        string
}

func (v Vehicle) String() string {
	var s string

	s = s + "Car - Brand: " + v.Brand
	s = s + ", Color: " + v.Color
	s = s + ", Max Speed: " + strconv.Itoa(v.Speed)
	s = s + ", Doors: " + strconv.Itoa(v.Doors)

	return s
}
