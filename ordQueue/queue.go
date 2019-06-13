package ordQueue

import (
	"github.com/infinivision/common/functor"
	"github.com/infinivision/common/typeclass"
)

func New() OrdQueue {
	return &ordQueueHead{
		queue: nil,
	}
}

// 将多个列表合成一个列表
func Concat(ls []OrdQueue) OrdQueue {
	xs := New()
	for i, j := 0, len(ls); i < j; i++ {
		ls[i].Foldl(func(x typeclass.Ord, b interface{}) interface{} {
			bs, _ := b.(OrdQueue)
			bs.Push(x)
			return bs
		}, xs)
	}
	return xs
}

func (xs *ordQueue) Prev() *ordQueue {
	return xs.prev
}

func (xs *ordQueue) Next() *ordQueue {
	return xs.next
}

// xs => xs -> ys
func (xs *ordQueue) AddNext(ys *ordQueue) *ordQueue {
	if !xs.Null() {
		xs.next = ys
	}
	return xs
}

// xs => ys -> xs
func (xs *ordQueue) AddPrev(ys *ordQueue) *ordQueue {
	if !xs.Null() {
		xs.prev = ys
	}
	return xs
}

// prev -> next => prev -> xs -> next
func Link(xs, prev, next *ordQueue) *ordQueue {
	xs.next = next
	xs.prev = prev
	prev.AddNext(xs)
	next.AddPrev(xs)
	return xs.Head()
}

// prev -> xs -> next => prev -> next
func UnLink(xs, prev, next *ordQueue) *ordQueue {
	switch {
	case prev.Null():
		next.AddPrev(prev)
		return next
	default:
		prev.AddNext(next)
		next.AddPrev(prev)
		return prev.Head()
	}
}

func (xs *ordQueue) Null() bool {
	return xs == nil
}

func (xs *ordQueue) Length() int {
	switch {
	case xs.Null():
		return 0
	default:
		return xs.Next().Length() + 1
	}
}

func (xs *ordQueue) Elem(x typeclass.Ord) typeclass.Ord {
	switch {
	case xs.Null():
		return nil
	case xs.node.Eq(x):
		return xs.node
	default:
		return xs.Next().Elem(x)
	}
}

func (xs *ordQueue) Map(f functor.MapFunc, ys OrdQueue) OrdQueue {
	switch {
	case xs.Null():
		return ys
	default:
		ys.Push(f(xs.node))
		return xs.Next().Map(f, ys)
	}
}

func (xs *ordQueue) Foldl(f functor.FoldFunc, b interface{}) interface{} {
	switch {
	case xs.Null():
		return b
	default:
		return xs.Next().Foldl(f, f(xs.node, b))
	}
}

func (xs *ordQueue) Foldr(f functor.FoldFunc, b interface{}) interface{} {
	switch {
	case xs.Null():
		return b
	default:
		return f(xs.node, xs.Next().Foldr(f, b))
	}
}

func (xs *ordQueue) Head() *ordQueue {
	switch {
	case xs.Null():
		return xs
	case xs.Prev().Null():
		return xs
	default:
		return xs.Prev().Head()
	}
}

func (xs *ordQueue) Last() *ordQueue {
	switch {
	case xs.Null():
		return xs
	case xs.Next().Null():
		return xs
	default:
		return xs.Next().Last()
	}
}

func (xs *ordQueue) Index(n int) *ordQueue {
	switch {
	case xs.Null():
		return xs
	case n == 0:
		return xs
	default:
		return xs.Next().Index(n - 1)
	}
}

func (xs *ordQueue) DeleteByIndex(n int) *ordQueue {
	switch {
	case xs.Null():
		return xs
	case n == 0:
		return UnLink(xs, xs.Prev(), xs.Next())
	default:
		return xs.Next().DeleteByIndex(n - 1)
	}
}

func (xs *ordQueue) Push(x typeclass.Ord) *ordQueue {
	switch {
	case xs.Null():
		return &ordQueue{node: x, next: nil, prev: nil}
	case xs.node.Ge(x):
		return Link(&ordQueue{node: x, next: nil, prev: nil}, xs.Prev(), xs)
	case xs.Next().Null():
		return Link(&ordQueue{node: x, next: nil, prev: nil}, xs, xs.Next())
	default:
		return xs.Next().Push(x)
	}
}

func (xs *ordQueue) Delete(x typeclass.Ord) *ordQueue {
	switch {
	case xs.Null(): // Empty List
		return xs
	case xs.node.Eq(x):
		return UnLink(xs, xs.Prev(), xs.Next())
	case xs.Next().Null(): // Can't Find x
		return xs.Head()
	default:
		return xs.Next().Delete(x)
	}
}

func (xs *ordQueue) ElemIndex(x typeclass.Ord) int {
	switch {
	case xs.Null():
		return -1
	case xs.node.Eq(x):
		return 0
	default:
		switch n := xs.Next().ElemIndex(x); n {
		case -1:
			return -1
		default:
			return n + 1
		}
	}
}

func (xs *ordQueue) DeleteBy(f func(typeclass.Ord) bool) *ordQueue {
	switch {
	case xs.Null(): // Empty List
		return xs
	case f(xs.node):
		return UnLink(xs, xs.Prev(), xs.Next())
	case xs.Next().Null(): // Can't Find x
		return xs.Head()
	default:
		return xs.Next().DeleteBy(f)
	}
}

func (xs *ordQueue) Filter(f func(typeclass.Ord) bool, ys OrdQueue) OrdQueue {
	switch {
	case xs.Null():
		return ys
	case f(xs.node):
		ys.Push(xs.node)
		return xs.Next().Filter(f, ys)
	default:
		return xs.Next().Filter(f, ys)
	}
}

func (xs *ordQueue) FoldWhile(f func(typeclass.Ord, interface{}) bool, b interface{}) interface{} {
	switch {
	case xs.Null():
		return b
	case !f(xs.node, b):
		return b
	default:
		return xs.Next().FoldWhile(f, b)
	}
}

func (xs *ordQueue) TakeWhile(f func(typeclass.Ord) bool, ys OrdQueue) OrdQueue {
	switch {
	case xs.Null():
		return ys
	case !f(xs.node):
		return ys
	default:
		ys.Push(xs.node)
		return xs.Next().TakeWhile(f, ys)
	}
}

func (xs *ordQueue) DropWhile(f func(typeclass.Ord) bool, ys OrdQueue) OrdQueue {
	switch {
	case xs.Null():
		return ys
	case !f(xs.node):
		xs.Foldl(func(x typeclass.Ord, b interface{}) interface{} {
			bs, _ := b.(OrdQueue)
			bs.Push(x)
			return bs
		}, ys)
		return ys
	default:
		return xs.Next().DropWhile(f, ys)
	}
}
