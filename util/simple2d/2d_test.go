package simple2d

import (
	"testing"
)

type dirtest struct {
	p, q, want Point
}

func TestDirectionTo(t *testing.T) {
	tests := []dirtest{
		dirtest{Point{}, Point{3, -3}, Point{1, -1}},
		dirtest{Point{}, Point{3, 3}, Point{1, 1}},
		dirtest{Point{}, Point{-3, 3}, Point{-1, 1}},
		dirtest{Point{}, Point{-3, -3}, Point{-1, -1}},
		dirtest{Point{3, -3}, Point{3, -3}, Point{0, 0}},
		dirtest{Point{-3, -3}, Point{3, -3}, Point{1, 0}},
		dirtest{Point{3, 3}, Point{3, -3}, Point{0, -1}},
	}

	for _, dt := range tests {
		ret := dt.p.DirectionTo(dt.q)
		if !ret.Equals(dt.want) {
			t.Fatalf("DirectionTo from %v to %v = %v, want %v", dt.p, dt.q, ret, dt.want)
		}
	}
}

func TestMHDistance(t *testing.T) {
	ret := Point{2, 3}.MhDistance(Point{7, 10})
	if ret != 12 {
		t.Fatalf("MhDistance from %v to %v = %d, want %d\n", Point{2, 3}, Point{7, 10}, ret, 12)
	}
}
