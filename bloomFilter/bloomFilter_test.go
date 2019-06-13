package bloomFilter

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	b, _ := New(3, 20000, 0)
	b.Insert([]byte("abc"))
	fmt.Printf("abc = %v\n", b.Elem([]byte("abc")))
	fmt.Printf("ab = %v\n", b.Elem([]byte("ab")))
}
