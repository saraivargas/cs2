package main

import (
	"fmt"

	"github.com/saraivargas/cs2/Assignment_3/vehicle"
)

func main() {
	var car vehicle.Vehicle
	var bike vehicle.Vehicle
	var truck vehicle.Vehicle

	car.Brand = "Toyota"
	car.Color = "Red"
	car.Speed = 180
	car.Doors = 4

	bike.Brand = "Trek"
	bike.Color = "Blue"
	bike.Speed = 60
	bike.Type = "Mountain"

	truck.Brand = "Ford"
	truck.Color = "Black"
	truck.Speed = 140
	truck.Load = "5000 lbs"

	fmt.Println(car.String())
}
