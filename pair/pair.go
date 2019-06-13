package pair

import (
	"reflect"

	"github.com/infinivision/common/typeclass"
)

// y == x
func Eq(x Pair) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.Eq)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.Eq)
				yy, _ := v.Value().(typeclass.Eq)
				return yy.Eq(xx)
			}
		}
		return false
	}
}

// y != x
func NotEq(x Pair) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.Eq)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.Eq)
				yy, _ := v.Value().(typeclass.Eq)
				return !yy.Eq(xx)
			}
		}
		return false
	}
}

// y < x
func Lt(x Pair) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.Ord)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.Ord)
				yy, _ := v.Value().(typeclass.Ord)
				return yy.Lt(xx)
			}
		}
		return false
	}
}

// y <= x
func Le(x Pair) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.Ord)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.Ord)
				yy, _ := v.Value().(typeclass.Ord)
				return yy.Le(xx)
			}
		}
		return false
	}
}

// y > x
func Gt(x Pair) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.Ord)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.Ord)
				yy, _ := v.Value().(typeclass.Ord)
				return yy.Gt(xx)
			}
		}
		return false
	}
}

// y > x
func Ge(x Pair) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.Ord)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.Ord)
				yy, _ := v.Value().(typeclass.Ord)
				return yy.Ge(xx)
			}
		}
		return false
	}
}

// y.value = y.value ++ x.value
func Suffix(x Pair) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdBytes)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.OrdBytes)
				yy, _ := v.Value().(typeclass.OrdBytes)
				return New(v.Key(), typeclass.NewBytes(append(yy.Bytes(), xx.Bytes()...)))
			}
		}
		return y
	}
}

// y.value = x.value ++ y.value
func Prefix(x Pair) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		v, _ := y.(Pair)
		if t := reflect.TypeOf(x.Value()); t == reflect.TypeOf(v.Value()) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdBytes)(nil)).Elem()):
				xx, _ := x.Value().(typeclass.OrdBytes)
				yy, _ := v.Value().(typeclass.OrdBytes)
				return New(v.Key(), typeclass.NewBytes(append(xx.Bytes(), yy.Bytes()...)))
			}
		}
		return y
	}
}
