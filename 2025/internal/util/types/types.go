package types

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

type Point3D struct {
	X, Y, Z int
}

func New3DPoint(x, y, z int) Point3D {
	return Point3D{
		x, y, z,
	}
}
