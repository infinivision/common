package bheap

import "github.com/infinivision/common/typeclass"

func New() *bheap {
	return &bheap{[]typeclass.Ord{}}
}
