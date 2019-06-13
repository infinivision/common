package pair

import "github.com/infinivision/common/typeclass"

type Pair interface {
	typeclass.Ord
	Key() typeclass.Ord
	Value() interface{}
	Update(interface{}) error
}

type pair struct {
	v interface{}
	k typeclass.Ord
}

func New(k typeclass.Ord, v interface{}) Pair {
	return &pair{k: k, v: v}
}

func (a *pair) Key() typeclass.Ord {
	return a.k
}

func (a *pair) Value() interface{} {
	return a.v
}

// update value of pair
func (a *pair) Update(v interface{}) error {
	a.v = v
	return nil
}

func (a *pair) Eq(b typeclass.Eq) bool {
	v, _ := b.(*pair)
	return a.k.Eq(v.k)
}

func (a *pair) NotEq(b typeclass.Eq) bool {
	return !a.Eq(b)
}

func (a *pair) Lt(b typeclass.Ord) bool {
	v, _ := b.(*pair)
	return a.k.Lt(v.k)
}

func (a *pair) Le(b typeclass.Ord) bool {
	v, _ := b.(*pair)
	return a.k.Le(v.k)
}

func (a *pair) Gt(b typeclass.Ord) bool {
	v, _ := b.(*pair)
	return a.k.Gt(v.k)
}

func (a *pair) Ge(b typeclass.Ord) bool {
	v, _ := b.(*pair)
	return a.k.Ge(v.k)
}

func (a *pair) Compare(b typeclass.Ord) typeclass.Ordering {
	switch {
	case a.Lt(b):
		return typeclass.LT
	case a.Eq(b):
		return typeclass.EQ
	default:
		return typeclass.GT
	}
}
