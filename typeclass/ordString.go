package typeclass

import "strings"

type ordString string

type OrdString interface {
	Ord
	String() string
}

func NewString(a string) OrdString {
	return ordString(a)
}

func MaxString(a, b OrdString) OrdString {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func MinString(a, b OrdString) OrdString {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func (a ordString) String() string {
	return string(a)
}

func (a ordString) Eq(b Eq) bool {
	v, _ := b.(ordString)
	return strings.Compare(string(a), string(v)) == 0
}

func (a ordString) NotEq(b Eq) bool {
	return !a.Eq(b)
}

func (a ordString) Lt(b Ord) bool {
	v, _ := b.(ordString)
	return strings.Compare(string(a), string(v)) < 0
}

func (a ordString) Le(b Ord) bool {
	v, _ := b.(ordString)
	return strings.Compare(string(a), string(v)) <= 0
}

func (a ordString) Gt(b Ord) bool {
	v, _ := b.(ordString)
	return strings.Compare(string(a), string(v)) > 0
}

func (a ordString) Ge(b Ord) bool {
	v, _ := b.(ordString)
	return strings.Compare(string(a), string(v)) >= 0
}

func (a ordString) Compare(b Ord) Ordering {
	switch {
	case a.Lt(b):
		return LT
	case a.Eq(b):
		return EQ
	default:
		return GT
	}
}
