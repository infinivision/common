package treeMap

import (
	"fmt"
	"testing"

	"github.com/infinivision/common/pair"
	"github.com/infinivision/common/typeclass"
)

func lPrint(x typeclass.Ord, n interface{}) interface{} {
	i, _ := n.(int)
	fmt.Printf("\t%d: %v\n", i, x)
	return i + 1
}

func Test(t *testing.T) {
	m := New()

	for i := 0; i < 100; i++ {
		m.Insert(typeclass.NewInt(i), typeclass.NewBytes([]byte{byte(i)}))
	}

	m.Foldl(lPrint, 0)

	m0 := m.Map(pair.Suffix(pair.New(nil, typeclass.NewBytes([]byte("xxx")))))

	m0.Foldl(lPrint, 0)

	m1 := m.Map(pair.Prefix(pair.New(nil, typeclass.NewBytes([]byte("xxx")))))
	m1.Foldl(lPrint, 0)

	m2 := m.Filter(pair.Lt(pair.New(nil, typeclass.NewBytes([]byte{byte(10)}))))
	m2.Foldl(lPrint, 0)
}
