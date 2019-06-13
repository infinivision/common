package slice

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/infinivision/common/curry"
	"github.com/infinivision/common/typeclass"
)

func Test(t *testing.T) {
	s := []typeclass.Ord{}
	s0 := []typeclass.Ord{}
	for i := 0; i < 100; i++ {
		s = append(s, typeclass.NewInt(rand.Intn(1000)))
	}
	fmt.Printf("Init: %v\n\n", s)

	s = Qsort(s)
	fmt.Printf("Qsort: %v\n\n", s)

	s0 = Nub(s)
	fmt.Printf("Nub: %v\n\n", s0)

	s0 = Filter(curry.Lt(typeclass.NewInt(728)), s)
	fmt.Printf("Filter: %v\n\n", s0)

	s0 = Push(typeclass.NewInt(730), s)
	fmt.Printf("Push: %v\n\n", s0)

	fmt.Printf("Elem 730: %v\n", Elem(typeclass.NewInt(730), s0))
	fmt.Printf("ElemIndex 730: %v\n", ElemIndex(typeclass.NewInt(730), s0))

	s0 = DeleteBy(curry.Eq(typeclass.NewInt(728)), s)
	fmt.Printf("Delete 728: %v\n\n", s0)

	s0 = Delete(typeclass.NewInt(730), s)
	fmt.Printf("Delete 730: %v\n\n", s0)

	s0 = Map(curry.Mul(typeclass.NewInt(10)), s)
	fmt.Printf("Map: %v\n\n", s0)

	fmt.Printf("Foldl: %v\n",
		Foldl(func(x typeclass.Ord, y interface{}) interface{} {
			xx, _ := x.(typeclass.OrdInt)
			yy, _ := y.(typeclass.OrdInt)
			return typeclass.NewInt(xx.Int() + yy.Int())
		}, typeclass.NewInt(0), s))
	fmt.Printf("Foldr: %v\n",
		Foldr(func(x typeclass.Ord, y interface{}) interface{} {
			xx, _ := x.(typeclass.OrdInt)
			yy, _ := y.(typeclass.OrdInt)
			return typeclass.NewInt(xx.Int() + yy.Int())
		}, typeclass.NewInt(0), s))
}
