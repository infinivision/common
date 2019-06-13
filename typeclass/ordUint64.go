package typeclass

type ordUint64 uint64

type OrdUint64 interface {
	Ord
	Uint64() uint64
}

func NewUint64(v uint64) OrdUint64 {
	return ordUint64(v)
}

func MaxUint64(a, b ordUint64) ordUint64 {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func MinUint64(a, b ordUint64) ordUint64 {
	switch {
	case a.Lt(b):
		return a
	default:
		return b
	}
}

func (a ordUint64) Uint64() uint64 {
	return uint64(a)
}

func (a ordUint64) Eq(b Eq) bool {
	v, _ := b.(ordUint64)
	return a == v
}

func (a ordUint64) NotEq(b Eq) bool {
	return !a.Eq(b)
}

func (a ordUint64) Lt(b Ord) bool {
	v, _ := b.(ordUint64)
	return a < v
}

func (a ordUint64) Le(b Ord) bool {
	v, _ := b.(ordUint64)
	return a <= v
}

func (a ordUint64) Gt(b Ord) bool {
	v, _ := b.(ordUint64)
	return a > v
}

func (a ordUint64) Ge(b Ord) bool {
	v, _ := b.(ordUint64)
	return a >= v
}

func (a ordUint64) Compare(b Ord) Ordering {
	switch {
	case a.Lt(b):
		return LT
	case a.Eq(b):
		return EQ
	default:
		return GT
	}
}
