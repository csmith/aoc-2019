package common

import (
	"fmt"
)

// Point represents a point on a 2D grid.
type Point struct {
	X, Y int64
}

// Plus returns a new Point which is the sum of this point and the other point.
func (p Point) Plus(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

// String returns a user-friendly string for debugging purposes.
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Manhattan calculates the Manhattan distance between this point and the other point.
func (p Point) Manhattan(other Point) int64 {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}
