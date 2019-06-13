package MurmurHash

import (
	"fmt"
	"testing"
)

func TestMurmurHash(t *testing.T) {
	msg := []byte("test")
	h := New(0)
	h.Write(msg)
	fmt.Printf("%v\n", h.Sum(nil))
}
