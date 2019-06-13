package rbtree

import (
	"github.com/infinivision/common/functor"
	"github.com/infinivision/common/miscellaneous"
	"github.com/infinivision/common/typeclass"
)

func New() RBtree {
	return &tree{tree: nil}
}

// left tree
func (t *treeNode) L() *treeNode {
	return t.left
}

// right tree
func (t *treeNode) R() *treeNode {
	return t.right
}

// is Black
func (t *treeNode) isB() bool {
	return t.Null() || t.color == B
}

// is Red
func (t *treeNode) isR() bool {
	return !t.Null() && t.color == R
}

// is double Black
func (t *treeNode) isBB() bool {
	return !t.Null() && t.color == BB
}

// is double Black Leaf
func (t *treeNode) isBBL() bool {
	return !t.Null() && t.color == BB && t.node == nil
}

// set node to Red
func (t *treeNode) setR() *treeNode {
	switch {
	case t.Null():
		return t
	default:
		t.color = R
		return t
	}
}

// set node to Black
func (t *treeNode) setB() *treeNode {
	switch {
	case t.isBBL():
		return nil
	case t.Null():
		return t
	default:
		t.color = B
		return t
	}
}

// set node to double Black
func (t *treeNode) setBB() *treeNode {
	switch {
	case t.Null():
		return &treeNode{color: BB, left: nil, right: nil, node: nil}
	default:
		t.color = BB
		return t
	}
}

// deepen
func (t *treeNode) setD() *treeNode {
	switch {
	case t.Null():
		return &treeNode{color: BB, left: nil, right: nil, node: nil}
	case t.isR():
		t.color = B
	case t.isB():
		t.color = BB
	}
	return t
}

// a = b
func (a *treeNode) set(b *treeNode) bool {
	switch {
	case b.Null():
		return false
	default:
		*a = *b
		return true
	}
}

func T(t, l, r *treeNode) *treeNode {
	t.left = l
	t.right = r
	return t
}

// balance tree
func balance(t, l, r *treeNode) *treeNode {
	a := new(treeNode)
	switch {
	case t.isB() && l.isR() && a.set(l.L()) && a.isR():
		return T(l.setR(), a.setB(), T(t.setB(), l.R(), r))
	case t.isB() && l.isR() && a.set(l.R()) && a.isR():
		return T(a.setR(), T(l.setB(), l.L(), a.L()), T(t.setB(), a.R(), r))
	case t.isB() && r.isR() && a.set(r.R()) && a.isR():
		return T(r.setR(), T(t.setB(), l, r.L()), a.setB())
	case t.isB() && r.isR() && a.set(r.L()) && a.isR():
		return T(a.setR(), T(t.setB(), l, a.L()), T(r.setB(), a.R(), r.R()))
	case t.isBB() && l.isR() && a.set(l.R()) && a.isR():
		return T(a.setB(), T(l.setB(), l.L(), a.L()), T(t.setB(), a.R(), r))
	case t.isBB() && r.isR() && a.set(r.L()) && a.isR():
		return T(a.setB(), T(t.setB(), t.L(), a.L()), T(r.setB(), a.R(), r.R()))
	default:
		return T(t, l, r)
	}
}

// bubble
func bubble(t, l, r *treeNode) *treeNode {
	a := new(treeNode)
	switch {
	case t.isB() && r.isBB() && l.isB():
		return balance(l.setBB(), l.L(), T(t.setR(), l.R(), r.setB()))
	case t.isB() && l.isBB() && r.isB():
		return balance(r.setBB(), T(t.setR(), l.setB(), r.L()), r.R())
	case t.isR() && l.isBB() && r.isB():
		return balance(r.setB(), T(t.setR(), l.setB(), r.L()), r.R())
	case t.isR() && r.isBB() && l.isB():
		return balance(l.setB(), l.L(), T(t.setR(), l.R(), r.setB()))
	case t.isB() && l.isBB() && r.isR() && a.set(r.L()) && a.isB():
		return balance(r.setB(), balance(t.setB(), l.setB(), a.setR()), r.R())
	case t.isB() && r.isBB() && l.isR() && a.set(l.R()) && a.isB():
		return balance(l.setB(), l.L(), balance(t.setB(), a.setR(), r.setB()))
	default:
		return T(t, l, r)
	}
}

func (t *treeNode) Null() bool {
	return t == nil
}

func (t *treeNode) Length() int {
	switch {
	case t.Null():
		return 0
	default:
		return t.L().Length() + t.R().Length() + 1
	}
}

func (t *treeNode) Elem(a typeclass.Ord) typeclass.Ord {
	switch {
	case t.Null():
		return nil
	case t.node.Eq(a):
		return t.node
	case t.node.Gt(a): // a < t
		return t.L().Elem(a)
	default:
		return t.R().Elem(a)
	}
}

func (t *treeNode) Minimum() *treeNode {
	switch {
	case t.Null():
		return nil
	case t.L().Null():
		return t
	default:
		return t.L().Minimum()
	}
}

func (t *treeNode) Maximum() *treeNode {
	switch {
	case t.Null():
		return nil
	case t.R().Null():
		return t
	default:
		return t.R().Maximum()
	}
}

func (t *treeNode) Map(f functor.MapFunc, u RBtree) RBtree {
	switch {
	case t.Null():
		return u
	default:
		u.Insert(f(t.node))
		t.L().Map(f, u)
		t.R().Map(f, u)
		return u
	}
}

func (t *treeNode) Foldl(f functor.FoldFunc, b interface{}) interface{} {
	switch {
	case t.Null():
		return b
	default:
		b = t.L().Foldl(f, b)
		b = f(t.node, b)
		b = t.R().Foldl(f, b)
		return b
	}
}

func (t *treeNode) Foldr(f functor.FoldFunc, b interface{}) interface{} {
	switch {
	case t.Null():
		return b
	default:
		b = t.R().Foldr(f, b)
		b = f(t.node, b)
		b = t.L().Foldr(f, b)
		return b
	}
}

func (t *treeNode) Root() typeclass.Ord {
	switch {
	case t.Null():
		return nil
	default:
		return t.node
	}
}

func (t *treeNode) Level() int {
	switch {
	case t.Null():
		return 0
	default:
		return miscellaneous.Max(t.L().Level(), t.R().Level()) + 1
	}
}

// 不支持插入重复的元素
func (t *treeNode) Insert(a typeclass.Ord) *treeNode {
	switch {
	case t.Null():
		return &treeNode{color: R, left: nil, right: nil, node: a}
	case a.Eq(t.node):
		return t
	case a.Lt(t.node): // a < t
		return balance(t, t.L().Insert(a), t.R())
	default:
		return balance(t, t.L(), t.R().Insert(a))
	}
}

func (t *treeNode) Delete(a typeclass.Ord) *treeNode {
	switch {
	case t.Null():
		return t
	case a.Lt(t.node) && !t.L().Null():
		return bubble(t, t.L().Delete(a), t.R())
	case a.Gt(t.node) && !t.R().Null():
		return bubble(t, t.L(), t.R().Delete(a))
	case a.Eq(t.node):
		switch {
		case t.L().Null():
			if t.isB() {
				return t.R().setD()
			}
			return t.R()
		case t.R().Null():
			if t.isB() {
				return t.L().setD()
			}
			return t.L()
		default:
			min := t.R().Minimum()
			r := t.R().Delete(min.node)
			switch {
			case t.isR():
				min.setR()
			default:
				min.setB()
			}
			return bubble(min, t.L(), r)
		}
	default:
		return t
	}
}

func (t *treeNode) Filter(f func(typeclass.Ord) bool, u RBtree) RBtree {
	switch {
	case t.Null():
		return u
	case f(t.node):
		u.Insert(t.node)
	}
	t.L().Filter(f, u)
	t.R().Filter(f, u)
	return u
}

func (t *treeNode) IsBalanced() bool {
	return t.isBlackSame() && t.isRedSep()
}

func (t *treeNode) isBlackSame() bool {
	s := t.blacks(0)
	for i, j := 1, len(s); i < j; i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func (t *treeNode) blacks(n int) []int {
	switch {
	case t.Null():
		return []int{n + 1}
	case t.isB():
		return append(t.L().blacks(n+1), t.R().blacks(n+1)...)
	default:
		return append(t.L().blacks(n), t.R().blacks(n)...)
	}
}

func (t *treeNode) isRedSep() bool {
	switch {
	case t.Null():
		return true
	case t.isR() && (t.L().isR() || t.R().isR()):
		return false
	default:
		return t.L().isRedSep() && t.R().isRedSep()
	}
}
