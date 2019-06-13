package rbtree

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/infinivision/common/curry"
	"github.com/infinivision/common/typeclass"
)

func lPrint(x typeclass.Ord, n interface{}) interface{} {
	i, _ := n.(int)
	fmt.Printf("\t%d: %v\n", i, x)
	return i + 1
}

func rPrint(x typeclass.Ord, n interface{}) interface{} {
	i, _ := n.(int)
	fmt.Printf("\t%d: %v\n", i, x)
	return i - 1
}

func Test(t *testing.T) {
	tree := New()

	fmt.Printf("isEmpty: %v\n", tree.Null())

	for i := 0; i < 100; i++ {
		tree.Insert(typeclass.NewInt(rand.Intn(1000)))
	}

	fmt.Printf("Init:\n")
	tree.Foldl(lPrint, 0)

	fmt.Printf("Max: %v\n", tree.Maximum())
	fmt.Printf("Min: %v\n", tree.Minimum())
	fmt.Printf("isEmpty: %v\n", tree.Null())
	fmt.Printf("Length: %v\n", tree.Length())
	fmt.Printf("Level: %v\n", tree.Level())
	fmt.Printf("Root: %v\n", tree.Root())

	v := typeclass.NewInt(828)
	fmt.Printf("%v is Exist: %v\n", v, tree.Elem(v))

	v = typeclass.NewInt(829)
	fmt.Printf("%v is Exist: %v\n", v, tree.Elem(v))

	v = typeclass.NewInt(828)
	tree.Delete(v)

	fmt.Printf("Delete 828:\n")
	tree.Foldr(rPrint, tree.Length()-1)

	tree0 := tree.Filter(curry.Lt(typeclass.NewInt(200)))

	fmt.Printf("Filter < 200:\n")
	tree0.Foldl(lPrint, 0)

	tree1 := tree.Map(curry.Mul(typeclass.NewInt(10)))
	fmt.Printf("MAP Mul 10:\n")
	tree1.Foldl(lPrint, 0)
}
