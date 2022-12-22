package simple2d

import (
	"math"

	sm "github.com/jmccarv/adventofcode/util/math"
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
	return Point{X: sm.Min(p.X, q.X), Y: sm.Min(p.Y, q.Y)}
}

func (p Point) Max(q Point) Point {
	return Point{X: sm.Max(p.X, q.X), Y: sm.Max(p.Y, q.Y)}
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) DirectionTo(q Point) Point {
	return Point{X: sm.Cmp(q.X, p.X), Y: sm.Cmp(q.Y, p.Y)}
}

// Manhattan distance
func (p Point) MhDistance(q Point) int {
	return sm.Abs(p.X-q.X) + sm.Abs(p.Y-q.Y)
}

// distance between two points, A & B:
// AB = √( (x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2 )
func (p Point) Distance(q Point) float64 {
	return math.Sqrt(math.Pow(float64(q.X-p.X), 2) + math.Pow(float64(q.Y-q.Y), 2))
}

type Box struct {
	TL, BR Point
}

func (b1 Box) Overlaps(b2 Box) bool {
	if b1.BR.X < b2.TL.X || b1.TL.X > b2.BR.X {
		return false
	}
	if b1.BR.Y < b2.TL.Y || b1.TL.Y > b2.BR.Y {
		return false
	}
	return true
}

func (b Box) Area() int {
	return (b.BR.Y - b.TL.Y) * (b.BR.X - b.TL.X)
}

func (b Box) NrPoints() int {
	return (b.BR.Y - b.TL.Y + 1) * (b.BR.X - b.TL.X + 1)
}
