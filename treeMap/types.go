package treeMap

import (
	"github.com/infinivision/common/functor"
	"github.com/infinivision/common/pair"
	"github.com/infinivision/common/rbtree"
	"github.com/infinivision/common/typeclass"
)

type TreeMap interface {
	Clear()
	Delete(typeclass.Ord) error
	Insert(typeclass.Ord, interface{}) error
	Filter(func(typeclass.Ord) bool) TreeMap
	functor.Functor
}

type treeMap struct {
	tree rbtree.RBtree
}

func (m *treeMap) Clear() {
	m.tree.Clear()
}

func (m *treeMap) Delete(k typeclass.Ord) error {
	return m.tree.Delete(pair.New(k, nil))
}

func (m *treeMap) Insert(k typeclass.Ord, v interface{}) error {
	if a := m.Elem(k); a != nil {
		p, _ := a.(pair.Pair)
		return p.Update(v)
	}
	return m.tree.Insert(pair.New(k, v))
}

func (m *treeMap) Filter(f func(typeclass.Ord) bool) TreeMap {
	return &treeMap{m.tree.Filter(f)}
}

func (m *treeMap) Null() bool {
	return m.tree.Null()
}

func (m *treeMap) Length() int {
	return m.tree.Length()
}

func (m *treeMap) Minimum() typeclass.Ord {
	return m.tree.Minimum()
}

func (m *treeMap) Maximum() typeclass.Ord {
	return m.tree.Maximum()
}

func (m *treeMap) Elem(k typeclass.Ord) typeclass.Ord {
	return m.tree.Elem(pair.New(k, nil))
}

func (m *treeMap) Map(f functor.MapFunc) functor.Functor {
	return m.tree.Map(f)
}

func (m *treeMap) Foldl(f functor.FoldFunc, b interface{}) interface{} {
	return m.tree.Foldl(f, b)
}

func (m *treeMap) Foldr(f functor.FoldFunc, b interface{}) interface{} {
	return m.tree.Foldr(f, b)
}
