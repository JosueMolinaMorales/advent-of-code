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
