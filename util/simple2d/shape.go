package simple2d

import (
	"fmt"

	sm "github.com/jmccarv/adventofcode/util/math"
)

// A shape is defined by a retangular area with points inside
// that rectangle set to define the shape.
// The shape is located in space by its top left corner (Loc)
type Shape struct {
	Size   Point              // Size of a box that would encompas this shape
	Loc    Point              // Where the top left of this shape lies in space
	Points map[Point]struct{} // which points are set within the bounds of this shape
}

type ShapeField struct {
	Bounds Box // use math.MinInt and math.MaxInt for an 'infinite' space
	Shapes []*Shape
}

func (sf ShapeField) DumpWindow(bounds Box) string {
	p := sf.AllPoints()
	ret := ""
	for y := bounds.TL.Y; y <= bounds.BR.Y; y++ {
		for x := bounds.TL.X; x <= bounds.BR.X; x++ {
			if _, ok := p[Point{x, y}]; ok {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}
	return ret
}

func (sf ShapeField) AllPoints() map[Point]struct{} {
	ret := make(map[Point]struct{})
	for _, s := range sf.Shapes {
		for p := range s.Points {
			ret[p.Add(s.Loc)] = struct{}{}
		}
	}
	return ret
}

func (sf *ShapeField) AddShape(s Shape) *Shape {
	if s.Size.X == 0 && s.Size.Y == 0 {
		return nil
	}
	sf.Shapes = append(sf.Shapes, &s)
	return &s
}

// dir is a direction, values for X and Y are expected to be -1, 0, or 1
// Returns true if movement succeeded without collision
// Returns false if the was a collision (either with a wall(boundry) or another shape)
// If a collision would occur, the movement it stopped just before the shape would collide.
func (sf ShapeField) StepShape(s *Shape, dir Point) bool {
	if dir.X == 0 && dir.Y == 0 {
		return true
	}

	n := *s
	n.Loc = n.Loc.Add(dir)

	// Check for wall collision first
	if !sf.Bounds.Contains(n.Bounds()) {
		return false
	}

	// Check collision with other shapes in this field
	for _, s2 := range sf.Shapes {
		if s2 != s && n.CollidesWith(*s2) {
			return false
		}
	}

	*s = n
	return true
}

func (s Shape) Bounds() Box {
	return Box{TL: s.Loc, BR: s.Loc.Add(s.Size).Sub(Point{1, 1})}
}

func (s1 Shape) CollidesWith(s2 Shape) bool {
	s1b := s1.Bounds()
	s2b := s2.Bounds()

	if !s1b.Overlaps(s2b) {
		// They can't collide if their bounding boxes don't overlap!
		return false
	}

	// Now we have to check individual points... seems like there should be a better way :)
	area := s1b
	if s1b.Area() > s2b.Area() {
		area = s2b
	}
	for x := area.TL.X; x <= area.BR.X; x++ {
		for y := area.TL.Y; y <= area.BR.Y; y++ {
			if s1.HasPoint(Point{x, y}) && s2.HasPoint(Point{x, y}) {
				return true
			}
		}
	}

	return false
}

func (s Shape) HasPoint(p Point) bool {
	_, ok := s.Points[p.Sub(s.Loc)]
	return ok
}

/*
func (s Shape) Translate(by Point) Shape {
	s.Loc = s.Loc.Add(by)
	return s
}
*/

/* parse an array of strings, one string per line, like this
 *    #
 *  #####
 *   ###
 *    #
 */
func NewShapeFromString(in []string, point byte) Shape {
	s := Shape{Size: Point{Y: len(in)}, Points: make(map[Point]struct{})}

	for _, l := range in {
		s.Size.X = sm.Max(s.Size.X, len(l))
	}

	for y, l := range in {
		for x := 0; x < len(l); x++ {
			if l[x] == point {
				s.Points[Point{X: x, Y: y}] = struct{}{}
			}
		}
	}
	return s
}

func NewShapeFromPoints(points ...Point) Shape {
	s := Shape{Points: make(map[Point]struct{})}

	if len(points) == 0 {
		return s
	}

	tl, br := points[0], points[0]
	for _, p := range points {
		tl = tl.Min(p)
		br = br.Max(p)
	}

	for _, p := range points {
		s.Points[p.Sub(tl)] = struct{}{}
	}

	s.Size = br.Sub(tl).Add(Point{1, 1})

	return s
}

func (s Shape) String() string {
	ret := fmt.Sprintf("At: %v  Size: %v\n", s.Loc, s.Size)
	for y := 0; y < s.Size.Y; y++ {
		for x := 0; x < s.Size.X; x++ {
			if _, ok := s.Points[Point{x, y}]; ok {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}
	return ret
}
