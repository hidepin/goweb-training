package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000

	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Printf("%v\n", dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Printf("%v\n", dilbert)

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Printf("%#v\n", w)
}
