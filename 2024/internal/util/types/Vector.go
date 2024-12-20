package types

import "fmt"

type Vector struct {
	X, Y int
}

func NewVector(cords ...int) *Vector {
	if len(cords) > 0 && len(cords) != 2 {
		panic("Coords must only be 2 values")
	}
	if len(cords) == 2 {
		return &Vector{X: cords[0], Y: cords[1]}
	}
	return &Vector{X: 0, Y: 0}
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}
