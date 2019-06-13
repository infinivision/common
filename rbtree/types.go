package rbtree

import (
	"github.com/infinivision/common/functor"
	"github.com/infinivision/common/typeclass"
)

/*
 * rbtree = rbtree{
 * 		left
 *		right
 * 		color, ordNode
 * } | Empty Node
 */

type Color int

const (
	R Color = iota
	B
	BB
)

type RBtree interface {
	Clear()
	Level() int                             // 获取树的高度
	IsBalanced() bool                       // 是否平衡?
	Root() typeclass.Ord                    // 获取树的根
	Insert(typeclass.Ord) error             // 插入一个新的元素
	Delete(typeclass.Ord) error             // 不重复删除
	Filter(func(typeclass.Ord) bool) RBtree // 从树中筛选出满足条件的元素，生成一颗新的树
	functor.Functor
}

type tree struct {
	tree *treeNode
}

type treeNode struct {
	color Color
	left  *treeNode
	right *treeNode
	node  typeclass.Ord
}

// functor.Null
func (t *tree) Null() bool {
	return t.tree.Null()
}

// functor.Length
func (t *tree) Length() int {
	return t.tree.Length()
}

// functor.Elem
func (t *tree) Elem(x typeclass.Ord) typeclass.Ord {
	return t.tree.Elem(x)
}

// functor.Minimum
func (t *tree) Minimum() typeclass.Ord {
	switch x := t.tree.Minimum(); x {
	case nil:
		return nil
	default:
		return x.node
	}
}

// functor.Maximum
func (t *tree) Maximum() typeclass.Ord {
	switch x := t.tree.Maximum(); x {
	case nil:
		return nil
	default:
		return x.node
	}
}

// functor.Map
func (t *tree) Map(f functor.MapFunc) functor.Functor {
	return t.tree.Map(f, New())
}

// functor.Foldl
func (t *tree) Foldl(f functor.FoldFunc, b interface{}) interface{} {
	return t.tree.Foldl(f, b)
}

// functor.Foldr
func (t *tree) Foldr(f functor.FoldFunc, b interface{}) interface{} {
	return t.tree.Foldr(f, b)
}

func (t *tree) Clear() {
	t.tree = nil
}

func (t *tree) Level() int {
	return t.tree.Level()
}

func (t *tree) IsBalanced() bool {
	return t.tree.IsBalanced()
}

func (t *tree) Root() typeclass.Ord {
	return t.tree.Root()
}

func (t *tree) Insert(a typeclass.Ord) error {
	t.tree = t.tree.Insert(a).setB()
	return nil
}

func (t *treeNode) M() *treeNode {
	switch {
	case t.isB() && t.L().isB() && t.R().isB():
		return T(t.setR(), t.L(), t.R())
	default:
		return t
	}
}

func (t *tree) Delete(a typeclass.Ord) error {
	t.tree = t.tree.M().Delete(a).setB()
	return nil
}

func (t *tree) Filter(f func(typeclass.Ord) bool) RBtree {
	return t.tree.Filter(f, New())
}
