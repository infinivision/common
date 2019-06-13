package ordQueue

import (
	"github.com/infinivision/common/functor"
	"github.com/infinivision/common/typeclass"
)

/*
 * ordQueue = ordQueue{
 * 		next
 * 		prev
 * 		ordNode
 * } | Empty Node
 */
type OrdQueue interface {
	Head() typeclass.Ord
	Last() typeclass.Ord
	Index(int) typeclass.Ord
	Push(typeclass.Ord) error
	Delete(typeclass.Ord) error
	ElemIndex(typeclass.Ord) int
	DeleteByIndex(int) typeclass.Ord
	DeleteBy(func(typeclass.Ord) bool) error
	Filter(func(typeclass.Ord) bool) OrdQueue
	TakeWhile(func(typeclass.Ord) bool) OrdQueue
	DropWhile(func(typeclass.Ord) bool) OrdQueue
	FoldWhile(func(typeclass.Ord, interface{}) bool, interface{}) interface{}
	functor.Functor
}

type ordQueueHead struct {
	queue *ordQueue
}

// 一个典型的队列定义
type ordQueue struct {
	prev *ordQueue
	next *ordQueue
	node typeclass.Ord
}

// functor.Null
func (xs *ordQueueHead) Null() bool {
	return xs.queue.Null()
}

// functor.Length
func (xs *ordQueueHead) Length() int {
	return xs.queue.Length()
}

// functor.Elem
func (xs *ordQueueHead) Elem(x typeclass.Ord) typeclass.Ord {
	return xs.queue.Elem(x)
}

// functor.Minimum
func (xs *ordQueueHead) Minimum() typeclass.Ord {
	return xs.Head()
}

// functor.Maximum
func (xs *ordQueueHead) Maximum() typeclass.Ord {
	return xs.Last()
}

// functor.Map
func (xs *ordQueueHead) Map(f functor.MapFunc) functor.Functor {
	return xs.queue.Map(f, New())
}

// functor.Foldl
func (xs *ordQueueHead) Foldl(f functor.FoldFunc, b interface{}) interface{} {
	return xs.queue.Foldl(f, b)
}

// functor.Foldlr
func (xs *ordQueueHead) Foldr(f functor.FoldFunc, b interface{}) interface{} {
	return xs.queue.Foldr(f, b)
}

// 获取列表的第一个元素
func (xs *ordQueueHead) Head() typeclass.Ord {
	switch x := xs.queue.Head(); x {
	case nil:
		return nil
	default:
		return x.node
	}
}

// 获取列表的最后一个元素
func (xs *ordQueueHead) Last() typeclass.Ord {
	switch x := xs.queue.Last(); x {
	case nil:
		return nil
	default:
		return x.node
	}
}

// 根据下标获取列表的元素
func (xs *ordQueueHead) Index(n int) typeclass.Ord {
	switch x := xs.queue.Index(n); x {
	case nil:
		return nil
	default:
		return x.node
	}
}

// 将x插入到一个有序列表xs
func (xs *ordQueueHead) Push(x typeclass.Ord) error {
	xs.queue = xs.queue.Push(x)
	return nil
}

// 删除第一个与x相等的元素
func (xs *ordQueueHead) Delete(x typeclass.Ord) error {
	xs.queue = xs.queue.Delete(x)
	return nil
}

func (xs *ordQueueHead) DeleteByIndex(n int) typeclass.Ord {
	x := xs.Index(n)
	xs.queue = xs.queue.DeleteByIndex(n)
	return x
}

// 查看x是否存在列表中，如果存在返回下标，不存在返回-1
func (xs *ordQueueHead) ElemIndex(x typeclass.Ord) int {
	return xs.queue.ElemIndex(x)
}

// 删除第一个满足要求的元素
func (xs *ordQueueHead) DeleteBy(f func(typeclass.Ord) bool) error {
	xs.queue = xs.queue.DeleteBy(f)
	return nil
}

// 从列表中筛选出满足条件的元素，生成一个新的列表
func (xs *ordQueueHead) Filter(f func(typeclass.Ord) bool) OrdQueue {
	return xs.queue.Filter(f, New())
}

// 从列表中不断的挑选满元素，直到碰到不满足f的元素为止
func (xs *ordQueueHead) TakeWhile(f func(typeclass.Ord) bool) OrdQueue {
	return xs.queue.TakeWhile(f, New())
}

// 从列表中不断的丢弃元素，直到碰到不满足f的元素为止
func (xs *ordQueueHead) DropWhile(f func(typeclass.Ord) bool) OrdQueue {
	return xs.queue.DropWhile(f, New())
}

func (xs *ordQueueHead) FoldWhile(f func(typeclass.Ord, interface{}) bool, b interface{}) interface{} {
	return xs.queue.FoldWhile(f, b)
}
