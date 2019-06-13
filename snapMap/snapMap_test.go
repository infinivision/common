package snapMap

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/base58"
)

func testPrint(k, v []byte, a interface{}) interface{} {
	i, _ := a.(int)
	fmt.Printf("\t%d: %v, %v\n", i, string(k), string(v))
	return i + 1
}

func Test(t *testing.T) {
	m := New()
	for i := 0; i < 10; i++ {
		m.Set([]byte(fmt.Sprintf("key_%v", i)),
			[]byte(fmt.Sprintf("data_%v", i)))
	}
	m.Set([]byte(fmt.Sprintf("key_%v", 0)),
		[]byte(fmt.Sprintf("data_%v", 3)))
	{
		data := m.Show()
		fmt.Printf("length: %v\n", len(data))
		hData := sha256.Sum256(data)
		fmt.Printf("hash: %v\n", base58.Encode(hData[:]))
	}
	m.Fold(testPrint, 0)

	m0 := New()
	for i := 0; i < 10; i++ {
		m0.Set([]byte(fmt.Sprintf("key_%v", i)),
			[]byte(fmt.Sprintf("data_%v", i)))
	}
	m0.Set([]byte(fmt.Sprintf("key_%v", 0)),
		[]byte(fmt.Sprintf("data_%v", 3)))
	data0 := m0.Show()
	{
		data := data0
		fmt.Printf("length: %v\n", len(data))
		hData := sha256.Sum256(data)
		fmt.Printf("hash: %v\n", base58.Encode(hData[:]))
	}
	m0.Fold(testPrint, 0)

	m1 := New()
	m1.Read(data0)
	{
		data := m1.Show()
		fmt.Printf("length: %v\n", len(data))
		hData := sha256.Sum256(data)
		fmt.Printf("hash: %v\n", base58.Encode(hData[:]))
	}

	fmt.Printf("Print:\n")
	m1.Fold(testPrint, 0)

}
