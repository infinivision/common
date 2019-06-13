package treeMap

import "github.com/infinivision/common/rbtree"

func New() TreeMap {
	return &treeMap{rbtree.New()}
}
