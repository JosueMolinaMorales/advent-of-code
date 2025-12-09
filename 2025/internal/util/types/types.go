package types

import "math"

type Point struct {
	Row int
	Col int
}

func NewPoint(row, col int) Point {
	return Point{
		Row: row,
		Col: col,
	}
}

func EuclideanDistance(p1, p2 Point) int {
	return int(math.Sqrt(math.Pow(float64(p2.Row)-float64(p1.Row), 2) + math.Pow(float64(p2.Col)-float64(p1.Col), 2)))
}

type Point3D struct {
	X, Y, Z int
}

func New3DPoint(x, y, z int) Point3D {
	return Point3D{
		x, y, z,
	}
}
