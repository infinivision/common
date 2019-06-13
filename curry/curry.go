package curry

import (
	"reflect"

	"github.com/infinivision/common/typeclass"
)

// y == x
func Eq(x typeclass.Ord) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		return y.Eq(x)
	}
}

// y != x
func NotEq(x typeclass.Ord) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		return !y.Eq(x)
	}
}

// y < x
func Lt(x typeclass.Ord) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		return y.Lt(x)
	}
}

// y <= x
func Le(x typeclass.Ord) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		return y.Le(x)
	}
}

// y > x
func Gt(x typeclass.Ord) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		return y.Gt(x)
	}
}

// y > x
func Ge(x typeclass.Ord) func(typeclass.Ord) bool {
	return func(y typeclass.Ord) bool {
		return y.Ge(x)
	}
}

// x * y
func Mul(x typeclass.Ord) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		if t := reflect.TypeOf(x); t == reflect.TypeOf(y) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdInt)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdInt)
				yy, _ := y.(typeclass.OrdInt)
				return typeclass.NewInt(xx.Int() * yy.Int())
			case t.Implements(reflect.TypeOf((*typeclass.OrdFloat)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdFloat)
				yy, _ := y.(typeclass.OrdFloat)
				return typeclass.NewFloat(xx.Float() * yy.Float())
			}
		}
		return y
	}
}

// y / x
func Div(x typeclass.Ord) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		if t := reflect.TypeOf(x); t == reflect.TypeOf(y) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdInt)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdInt)
				yy, _ := y.(typeclass.OrdInt)
				return typeclass.NewInt(yy.Int() / xx.Int())
			case t.Implements(reflect.TypeOf((*typeclass.OrdFloat)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdFloat)
				yy, _ := y.(typeclass.OrdFloat)
				return typeclass.NewFloat(yy.Float() / xx.Float())
			}
		}
		return y
	}
}

// x + y
func Add(x typeclass.Ord) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		if t := reflect.TypeOf(x); t == reflect.TypeOf(y) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdInt)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdInt)
				yy, _ := y.(typeclass.OrdInt)
				return typeclass.NewInt(xx.Int() + yy.Int())
			case t.Implements(reflect.TypeOf((*typeclass.OrdFloat)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdFloat)
				yy, _ := y.(typeclass.OrdFloat)
				return typeclass.NewFloat(xx.Float() + yy.Float())
			}
		}
		return y
	}
}

// y - x
func Sub(x typeclass.Ord) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		if t := reflect.TypeOf(x); t == reflect.TypeOf(y) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdInt)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdInt)
				yy, _ := y.(typeclass.OrdInt)
				return typeclass.NewInt(yy.Int() - xx.Int())
			case t.Implements(reflect.TypeOf((*typeclass.OrdFloat)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdFloat)
				yy, _ := y.(typeclass.OrdFloat)
				return typeclass.NewFloat(yy.Float() - xx.Float())
			}
		}
		return y
	}
}

// y ++ x
func Suffix(x typeclass.Ord) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		if t := reflect.TypeOf(x); t == reflect.TypeOf(y) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdBytes)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdBytes)
				yy, _ := y.(typeclass.OrdBytes)
				return typeclass.NewBytes(append(yy.Bytes(), xx.Bytes()...))
			}
		}
		return y
	}
}

// x ++ y
func Prefix(x typeclass.Ord) func(typeclass.Ord) typeclass.Ord {
	return func(y typeclass.Ord) typeclass.Ord {
		if t := reflect.TypeOf(x); t == reflect.TypeOf(y) {
			switch {
			case t.Implements(reflect.TypeOf((*typeclass.OrdBytes)(nil)).Elem()):
				xx, _ := x.(typeclass.OrdBytes)
				yy, _ := y.(typeclass.OrdBytes)
				return typeclass.NewBytes(append(xx.Bytes(), yy.Bytes()...))
			}
		}
		return y
	}
}
