package trie

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func testPrint(k, v []byte, a interface{}) interface{} {
	i, _ := a.(int)
	fmt.Printf("\t%d: %v, %v\n", i, string(k), string(v))
	return i + 1
}

func Test(t *testing.T) {
	tr := New()
	for i := 0; i < 10; i++ {
		tr.Set([]byte(fmt.Sprintf("key_%v", i)),
			[]byte(fmt.Sprintf("data_%v", i)))
	}
	/*
		fmt.Printf("Print: \n")
		tr.Fold(testPrint, 0)

		data := tr.Show()
		fmt.Printf("data: %v\n", len(data))
		tr0 := New()
		_, err := tr0.Read(data)
		fmt.Printf("read: %v\n", err)
		fmt.Printf("Print: \n")
		tr0.Fold(testPrint, 0)
	*/
	d0 := sha256.Sum256(tr.Show())
	d1 := sha256.Sum256(tr.Show())
	fmt.Printf("%v\n", d0)
	fmt.Printf("%v\n", d1)
}
