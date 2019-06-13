package typeclass

type ordInt int

type OrdInt interface {
	Ord
	Int() int
}

func NewInt(v int) OrdInt {
	return ordInt(v)
}

func MaxInt(a, b ordInt) ordInt {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func MinInt(a, b ordInt) ordInt {
	switch {
	case a.Lt(b):
		return a
	default:
		return b
	}
}

func (a ordInt) Int() int {
	return int(a)
}

func (a ordInt) Eq(b Eq) bool {
	v, _ := b.(ordInt)
	return a == v
}

func (a ordInt) NotEq(b Eq) bool {
	return !a.Eq(b)
}

func (a ordInt) Lt(b Ord) bool {
	v, _ := b.(ordInt)
	return a < v
}

func (a ordInt) Le(b Ord) bool {
	v, _ := b.(ordInt)
	return a <= v
}

func (a ordInt) Gt(b Ord) bool {
	v, _ := b.(ordInt)
	return a > v
}

func (a ordInt) Ge(b Ord) bool {
	v, _ := b.(ordInt)
	return a >= v
}

func (a ordInt) Compare(b Ord) Ordering {
	switch {
	case a.Lt(b):
		return LT
	case a.Eq(b):
		return EQ
	default:
		return GT
	}
}
