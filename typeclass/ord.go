package typeclass

type Ordering int

const (
	LT Ordering = iota
	EQ
	GT
)

type Ord interface {
	Eq
	Lt(Ord) bool // <
	Le(Ord) bool // <=
	Gt(Ord) bool // >
	Ge(Ord) bool // >=
	Compare(Ord) Ordering
}

func Min(a, b Ord) Ord {
	switch {
	case a.Lt(b):
		return a
	default:
		return b
	}
}

func Max(a, b Ord) Ord {
	switch {
	case a.Lt(b):
		return b
	default:
		return a
	}
}
