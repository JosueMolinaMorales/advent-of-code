package types

import (
	"fmt"
	"math"
)

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

func (v Vector) ManhanttanDistance(o Vector) int {
	return int(math.Abs(float64(v.X-o.X)) + math.Abs(float64(v.Y-o.Y)))
}
