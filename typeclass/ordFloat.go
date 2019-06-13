package typeclass

type ordFloat float64

type OrdFloat interface {
	Ord
	Float() float64
}

func NewFloat(v float64) ordFloat {
	return ordFloat(v)
}

func MaxFloat(a, b ordFloat) ordFloat {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}

func MinFloat(a, b ordFloat) ordFloat {
	switch {
	case a.Lt(b):
		return a
	default:
		return b
	}
}

func (a ordFloat) Float() float64 {
	return float64(a)
}

func (a ordFloat) Eq(b Eq) bool {
	v, _ := b.(ordFloat)
	return a == v
}

func (a ordFloat) NotEq(b Eq) bool {
	return !a.Eq(b)
}

func (a ordFloat) Lt(b Ord) bool {
	v, _ := b.(ordFloat)
	return a < v
}

func (a ordFloat) Le(b Ord) bool {
	v, _ := b.(ordFloat)
	return a <= v
}

func (a ordFloat) Gt(b Ord) bool {
	v, _ := b.(ordFloat)
	return a > v
}

func (a ordFloat) Ge(b Ord) bool {
	v, _ := b.(ordFloat)
	return a >= v
}

func (a ordFloat) Compare(b Ord) Ordering {
	switch {
	case a.Lt(b):
		return LT
	case a.Eq(b):
		return EQ
	default:
		return GT
	}
}
