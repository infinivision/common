package bloomFilter

import "github.com/infinivision/common/typeclass"

const (
	M_OFF  = 13
	M_MASK = 0x1FFF
)

type BloomFilter interface {
	Elem([]byte) bool
	Insert([]byte) error
	typeclass.Read
	typeclass.Show
}

type bloomFilter struct {
	k    uint32 // number of hash function
	m    uint32 // size of Bloom Filter in Byte
	data []byte
}
