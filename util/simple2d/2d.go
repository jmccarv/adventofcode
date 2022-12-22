package simple2d

import (
	"github.com/jmccarv/adventofcode/util/math"
)

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}

func (p Point) Min(q Point) Point {
	return Point{X: math.Min(p.X, q.X), Y: math.Min(p.Y, q.Y)}
}

func (p Point) Max(q Point) Point {
	return Point{X: math.Max(p.X, q.X), Y: math.Max(p.Y, q.Y)}
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) DirectionTo(q Point) Point {
	return Point{X: math.Cmp(q.X, p.X), Y: math.Cmp(q.Y, p.Y)}
}

type Box struct {
	TL, BR Point
}
