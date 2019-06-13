package typeclass

import "bytes"

type ordBytes []byte

type OrdBytes interface {
	Ord
	Bytes() []byte
}

func NewBytes(a []byte) OrdBytes {
	return ordBytes(a)
}

func MaxBytes(a, b OrdBytes) OrdBytes {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func MinBytes(a, b OrdBytes) OrdBytes {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func (a ordBytes) Bytes() []byte {
	return []byte(a)
}

func (a ordBytes) Eq(b Eq) bool {
	v, _ := b.(ordBytes)
	return bytes.Compare([]byte(a), []byte(v)) == 0
}

func (a ordBytes) NotEq(b Eq) bool {
	return !a.Eq(b)
}

func (a ordBytes) Lt(b Ord) bool {
	v, _ := b.(ordBytes)
	return bytes.Compare([]byte(a), []byte(v)) < 0
}

func (a ordBytes) Le(b Ord) bool {
	v, _ := b.(ordBytes)
	return bytes.Compare([]byte(a), []byte(v)) <= 0
}

func (a ordBytes) Gt(b Ord) bool {
	v, _ := b.(ordBytes)
	return bytes.Compare([]byte(a), []byte(v)) > 0
}

func (a ordBytes) Ge(b Ord) bool {
	v, _ := b.(ordBytes)
	return bytes.Compare([]byte(a), []byte(v)) >= 0
}

func (a ordBytes) Compare(b Ord) Ordering {
	switch {
	case a.Lt(b):
		return LT
	case a.Eq(b):
		return EQ
	default:
		return GT
	}
}
