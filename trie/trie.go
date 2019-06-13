package trie

import "github.com/infinivision/common/miscellaneous"

func New() *tree {
	return &tree{tree: &TreeNode{nil, make(map[uint8]*TreeNode)}}
}

func (t *TreeNode) Get(k []byte) []byte {
	switch {
	case t == nil:
		return nil
	case len(k) == 0:
		return t.Value
	default:
		return t.getNode(uint8(k[0])).Get(k[1:])
	}
}

func (t *TreeNode) Set(k, v []byte) error {
	switch {
	case len(k) == 0:
		t.Value = miscellaneous.Dup(v)
		return nil
	default:
		return t.getNodeOrNew(uint8(k[0])).Set(k[1:], v)
	}
}

func (t *TreeNode) Fold(p []byte, f FoldFunc, b interface{}) interface{} {
	switch {
	case t == nil:
		return b
	default:
		if t.Value != nil {
			b = f(p, t.Value, b)
		}
		for i := 0; i < CHILDREN; i++ {
			b = t.getNode(uint8(i)).Fold(append(p, byte(i)), f, b)
		}
		return b
	}
}

func (t *TreeNode) getNode(k uint8) *TreeNode {
	if _, ok := t.Children[k]; !ok {
		return nil
	}
	return t.Children[k]
}

func (t *TreeNode) getNodeOrNew(k uint8) *TreeNode {
	if _, ok := t.Children[k]; !ok {
		t.Children[k] = &TreeNode{nil, make(map[uint8]*TreeNode)}
	}
	return t.Children[k]
}
