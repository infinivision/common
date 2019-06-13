/*
 * 因为golang对尾递归不做优化处理，所以只能改变编码的方式
 * 关于golang是否应该优化尾递归的讨论可以参见: https://github.com/golang/go/issues/16798
 */
package slice

import (
	"github.com/infinivision/common/curry"
	"github.com/infinivision/common/functor"
	"github.com/infinivision/common/typeclass"
)

/*
func Nub(xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	default:
		return append([]typeclass.Ord{xs[0]}, Nub(Filter(curry.NotEq(xs[0]), xs[1:]))...)
	}
}
*/
func Nub(xs []typeclass.Ord) []typeclass.Ord {
	var rs []typeclass.Ord

	for {
		switch {
		case len(xs) == 0:
			return rs
		default:
			rs = append(rs, xs[0])
			xs = Filter(curry.NotEq(xs[0]), xs[1:])
		}
	}
}

// 从数组xs中提取满足f要求的元素
/*
func Filter(f func(typeclass.Ord) bool, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	case f(xs[0]):
		return append([]typeclass.Ord{xs[0]}, Filter(f, xs[1:])...)
	default:
		return Filter(f, xs[1:])
	}
}
*/
func Filter(f func(typeclass.Ord) bool, xs []typeclass.Ord) []typeclass.Ord {
	var rs []typeclass.Ord

	for {
		switch {
		case len(xs) == 0:
			return rs
		case f(xs[0]):
			rs = append(rs, xs[0])
		}
		xs = xs[1:]
	}
}

// 判断数组中是否存在某个元素
/*
func Elem(x typeclass.Ord, xs []typeclass.Ord) typeclass.Ord {
	switch {
	case len(xs) == 0:
		return nil
	case xs[0].Eq(x):
		return xs[0]
	default:
		return Elem(x, xs[1:])
	}
}
*/
func Elem(x typeclass.Ord, xs []typeclass.Ord) typeclass.Ord {
	for {
		switch {
		case len(xs) == 0:
			return nil
		case xs[0].Eq(x):
			return xs[0]
		default:
			xs = xs[1:]
		}
	}
}

// 判断数组中是否存在某个元素，存在返回下标，不存在返回-1
/*
func ElemIndex(x typeclass.Ord, xs []typeclass.Ord) int {
	switch {
	case len(xs) == 0:
		return -1
	case xs[0].Eq(x):
		return 0
	default:
		switch n := ElemIndex(x, xs[1:]); n {
		case -1:
			return -1
		default:
			return n + 1
		}
	}
}
*/
func ElemIndex(x typeclass.Ord, xs []typeclass.Ord) int {
	for i, j := 0, len(xs); i < j; i++ {
		if xs[i].Eq(x) {
			return i
		}
	}
	return -1
}

// 删除数组中的元素x，不重复删除
/*
func Delete(x typeclass.Ord, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	case xs[0].Eq(x):
		return xs[1:]
	default:
		return append([]typeclass.Ord{xs[0]}, Delete(x, xs[1:])...)
	}
}
*/
func Delete(x typeclass.Ord, xs []typeclass.Ord) []typeclass.Ord {
	var rs []typeclass.Ord

	for {
		switch {
		case len(xs) == 0:
			return rs
		case xs[0].Eq(x):
			return append(rs, xs[1:]...)
		default:
			rs = append(rs, xs[0])
			xs = xs[1:]
		}
	}
}

// 删除数组xs中满足f条件的元素，不重复删除
/*
func DeleteBy(f func(typeclass.Ord) bool, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	case f(xs[0]):
		return xs[1:]
	default:
		return append([]typeclass.Ord{xs[0]}, DeleteBy(f, xs[1:])...)
	}
}
*/
func DeleteBy(f func(typeclass.Ord) bool, xs []typeclass.Ord) []typeclass.Ord {
	var rs []typeclass.Ord

	for {
		switch {
		case len(xs) == 0:
			return rs
		case f(xs[0]):
			return append(rs, xs[1:]...)
		default:
			rs = append(rs, xs[0])
			xs = xs[1:]
		}
	}
}

/*
func Qsort(xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	default:
		return append(append(Qsort(Filter(curry.Lt(xs[0]), xs[1:])), xs[0]),
			Qsort(Filter(curry.Ge(xs[0]), xs[1:]))...)
	}
}
*/

type sliceRange struct {
	x int
	y int
}

func Qsort(xs []typeclass.Ord) []typeclass.Ord {
	var qs []*sliceRange

	ys := make([]typeclass.Ord, len(xs))
	copy(ys, xs)
	qs = append(qs, &sliceRange{0, len(ys) - 1})
	for len(qs) > 0 {
		x, y := qs[0].x, qs[0].y
		if x < y {
			z := ys[x]
			ls := Filter(curry.Lt(z), ys[x+1:y+1])
			rs := Filter(curry.Ge(z), ys[x+1:y+1])
			if len(ls) != 0 {
				copy(ys[x:x+len(ls)], ls)
				qs = append(qs, &sliceRange{x, x + len(ls) - 1})
			}
			if len(rs) != 0 {
				copy(ys[x+len(ls)+1:y+1], rs)
				qs = append(qs, &sliceRange{x + len(ls) + 1, y})
			}
			ys[x+len(ls)] = z
		}
		qs = qs[1:]
	}
	return ys
}

/*
func Bsearch(xs []typeclass.Ord, x typeclass.Ord) bool {
	switch {
	case len(xs) == 0:
		return false
	case len(xs) == 1:
		return xs[0] == x
	case x.Eq(xs[len(xs)/2]):
		return true
	case x.Gt(xs[len(xs)/2]):
		return Bsearch(xs[len(xs)/2:], x)
	default:
		return Bsearch(xs[:len(xs)/2], x)
	}
}
*/
func Bsearch(x typeclass.Ord, xs []typeclass.Ord) int {
	mid, start, end := 0, 0, len(xs)-1
	for start <= end {
		mid = start + (end-start)/2
		switch {
		case xs[mid].Eq(x):
			return mid
		case xs[mid].Lt(x):
			start = mid + 1
		default:
			end = mid - 1
		}
	}
	return -1
}

// 假设xs是个有序队列(从小到大)，然后insert(x)后满足x0 <= x <= x1
/*
func Push(x typeclass.Ord, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return []typeclass.Ord{x}
	case xs[0].Ge(x):
		return append([]typeclass.Ord{x}, xs...)
	default:
		return append(xs[:1], Push(x, xs[1:])...)
	}
}
*/
func Push(x typeclass.Ord, xs []typeclass.Ord) []typeclass.Ord {
	var rs []typeclass.Ord

	for {
		switch {
		case len(xs) == 0:
			return append(rs, x)
		case xs[0].Ge(x):
			rs = append(rs, x)
			return append(rs, xs...)
		default:
			rs = append(rs, xs[0])
			xs = xs[1:]
		}
	}
}

// 将一个列表根据f映射成另一个元素
func Map(f functor.MapFunc, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	default:
		return append([]typeclass.Ord{f(xs[0])}, Map(f, xs[1:])...)
	}
}

// 将一个列表根据x和f从左到右折叠成一个新的元素
func Foldl(f functor.FoldFunc, x interface{}, xs []typeclass.Ord) interface{} {
	switch {
	case len(xs) == 0:
		return x
	default:
		return Foldl(f, f(xs[0], x), xs[1:])
	}
}

// 将一个列表根据x和f从右到左折叠成一个新的元素
func Foldr(f functor.FoldFunc, x interface{}, xs []typeclass.Ord) interface{} {
	switch {
	case len(xs) == 0:
		return x
	default:
		return f(xs[0], Foldr(f, x, xs[1:]))
	}
}

func FoldWhile(f func(typeclass.Ord, interface{}) bool, x interface{}, xs []typeclass.Ord) interface{} {
	switch {
	case len(xs) == 0:
		return x
	case !f(xs[0], x):
		return x
	default:
		return FoldWhile(f, f(xs[0], x), xs[1:])
	}
}
