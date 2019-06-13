package ordQueue

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/infinivision/common/curry"
	"github.com/infinivision/common/typeclass"
)

type blockInfo struct {
	index int
}

func lPrint(x typeclass.Ord, n interface{}) interface{} {
	i, _ := n.(int)
	fmt.Printf("%d: %v\n", i, x)
	return i + 1
}

func rPrint(x typeclass.Ord, n interface{}) interface{} {
	i, _ := n.(int)
	fmt.Printf("%d: %v\n", i, x)
	return i - 1
}

func lPrint0(x typeclass.Ord, n interface{}) bool {
	bi, _ := n.(*blockInfo)
	defer func() {
		bi.index = bi.index + 1
	}()
	fmt.Printf("%d: %v\n", bi.index, x)
	if bi.index < 50 {
		return true
	} else {
		return false
	}
}

func Test(t *testing.T) {
	queue := New()

	fmt.Printf("isEmpty %v\n", queue.Null())

	for i := 100; i >= 0; i-- {
		queue.Push(typeclass.NewInt(rand.Intn(1000)))
	}

	queue.Foldl(lPrint, 0)
	fmt.Printf("10 = %v\n", queue.Index(10))
	fmt.Printf("100 = %v\n", queue.Index(100))
	fmt.Printf("1000 = %v\n", queue.Index(1000))
	fmt.Printf("head = %v\n", queue.Head())
	fmt.Printf("last = %v\n", queue.Last())

	fmt.Printf("isEmpty %v\n", queue.Null())
	fmt.Printf("Length = %v\n", queue.Length())

	v := typeclass.NewInt(828)

	fmt.Printf("%v is Exit: %v\n", v, queue.Elem(v))
	fmt.Printf("%v index is Exit: %v\n", v, queue.ElemIndex(v))

	v = typeclass.NewInt(829)
	fmt.Printf("%v is Exit: %v\n", v, queue.Elem(v))
	fmt.Printf("%v index is Exit: %v\n", v, queue.ElemIndex(v))

	v = typeclass.NewInt(828)
	queue.Delete(v)

	fmt.Printf("Delete 828:\n")
	queue.Foldr(rPrint, queue.Length()-1)

	queue0 := queue.Filter(curry.Lt(typeclass.NewInt(800)))
	fmt.Printf("Filter < 800:\n")
	queue0.Foldl(lPrint, 0)

	queue1 := queue.TakeWhile(curry.Lt(typeclass.NewInt(300)))
	fmt.Printf("TakeWhile < 300:\n")
	queue1.Foldl(lPrint, 0)

	queue2 := queue.DropWhile(curry.Lt(typeclass.NewInt(400)))
	fmt.Printf("DropWhile < 400:\n")
	queue2.Foldl(lPrint, 0)

	queue3 := Concat([]OrdQueue{queue1, queue2})
	fmt.Printf("Concat:\n")
	queue3.Foldl(lPrint, 0)

	queue.DeleteBy(curry.Gt(typeclass.NewInt(200)))
	fmt.Printf("Delete > 200:\n")
	queue.Foldl(lPrint, 0)

	queue4 := queue.Map(curry.Mul(typeclass.NewInt(10)))
	fmt.Printf("MAP:\n")
	queue4.Foldl(lPrint, 0)

	fmt.Printf("Minimum = %v\n", queue4.Minimum())
	fmt.Printf("Maximum = %v\n", queue4.Maximum())
}
