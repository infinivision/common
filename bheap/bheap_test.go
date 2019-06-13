package bheap

import (
	"fmt"
	"testing"

	"github.com/infinivision/common/typeclass"
)

func Test(t *testing.T) {
	h := New()
	for i := 20; i > 0; i-- {
		h.Insert(typeclass.NewInt(i))
	}
	for i := 0; i < 10; i = i + 2 {
		h.Insert(typeclass.NewInt(i))
	}
	for x := h.Extract(); x != nil; x = h.Extract() {
		fmt.Printf("%v\n", x)
	}
}
