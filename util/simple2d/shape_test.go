package simple2d

import (
	"testing"
)

func TestParse(t *testing.T) {
	s := NewShapeFromString([]string{
		"  ### ",
		" #####",
		"   #",
		"",
		"##",
	}, '#')

	if !s.Size.Equals(Point{6, 5}) {
		t.Fatalf("NewShapeFromString size = %v, want %v\n", s.Size, Point{6, 5})
	}

	if len(s.Points) != 11 {
		t.Fatalf("NewShapeFromString nrPoints = %v, want = 11\n", len(s.Points))
	}

	s.Loc = Point{3, 4}
	b := s.Bounds()
	want := Box{Point{3, 4}, Point{8, 8}}
	if !b.Equals(want) {
		t.Fatalf("NewShapeFromString bounds = %v, want %v\n", b, want)
	}
}

func TestCollision(t *testing.T) {
	s1 := NewShapeFromString([]string{
		"  ### ",
		" ## ##",
		"  ###",
	}, '#')

	s2 := NewShapeFromString([]string{
		".",
	}, '.')

	if s1.CollidesWith(s2) {
		t.Fatalf("CollidesWith - These collide but should not:\n%v\n%v\n", s1, s2)
	}

	s2.Loc = Point{1, 1}
	if !s1.CollidesWith(s2) {
		t.Fatalf("CollidesWith - These should collide but don't:\n%v\n%v\n", s1, s2)
	}

	s2.Loc = Point{3, 1}
	if s1.CollidesWith(s2) {
		t.Fatalf("CollidesWith - These collide but should not:\n%v\n%v\n", s1, s2)
	}

	s2.Loc = Point{-10, 1}
	if s1.CollidesWith(s2) {
		t.Fatalf("CollidesWith - These collide but should not:\n%v\n%v\n", s1, s2)
	}
}
