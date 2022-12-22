package math

import (
	"golang.org/x/exp/constraints"
)

type Signed interface {
	constraints.Signed | constraints.Float
}

func Max[T constraints.Ordered](a ...T) T {
	if len(a) < 1 {
		return T(0)
	}
	ret := a[0]
	for _, val := range a[1:] {
		if val > ret {
			ret = val
		}
	}
	return ret
}

func Min[T constraints.Ordered](a ...T) T {
	if len(a) < 1 {
		return T(0)
	}
	ret := a[0]
	for _, val := range a[1:] {
		if val < ret {
			ret = val
		}
	}
	return ret
}

func Cmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func Sign[T Signed](a T) T {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -1
	}
	return 1
}

func Abs[T Signed](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
