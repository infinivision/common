package bheap

import (
	"container/heap"

	"github.com/infinivision/common/typeclass"
)

type BHeap interface {
	Len() int
	Insert(typeclass.Ord)
	Extract() typeclass.Ord
}

type bheap struct {
	xs []typeclass.Ord
}

func (a *bheap) Len() int {
	return len(a.xs)
}

func (a *bheap) Swap(i, j int) {
	a.xs[i], a.xs[j] = a.xs[j], a.xs[i]
}

func (a *bheap) Less(i, j int) bool {
	return a.xs[i].Lt(a.xs[j])
}

func (a *bheap) Push(b interface{}) {
	x, _ := b.(typeclass.Ord)
	a.xs = append(a.xs, x)
}

func (a *bheap) Pop() interface{} {
	x := a.xs[len(a.xs)-1]
	a.xs = a.xs[:len(a.xs)-1]
	return x
}

func (a *bheap) Insert(b typeclass.Ord) {
	heap.Push(a, b)
}

func (a *bheap) Extract() typeclass.Ord {
	switch {
	case len(a.xs) == 0:
		return nil
	default:
		x := heap.Pop(a)
		return x.(typeclass.Ord)
	}
}
