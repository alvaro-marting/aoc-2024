package pkg

type Coord struct {
	X, Y int
}

func (c Coord) NextCoord(c1 Coord) Coord {
	return Coord{
		X: c.X + c1.X,
		Y: c.Y + c1.Y,
	}
}

var Axii = []Coord{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

var Diagonals = []Coord{
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

var Directions = []Coord{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}
