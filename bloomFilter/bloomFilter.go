package bloomFilter

import (
	"math"

	"github.com/infinivision/common/hash/MurmurHash"
	"github.com/infinivision/common/miscellaneous"
)

func (b *bloomFilter) getbit(v uint32) bool {
	x := v & (b.m - 1)
	y := v & 0x7
	return b.data[x]&(1<<y) != 0
}

func (b *bloomFilter) setbit(v uint32) {
	x := v & (b.m - 1)
	y := v & 0x7
	b.data[x] |= 1 << y
}

func genProbability(k, n uint32) (float64, float64, float64) {
	p := float64(0.1)
	switch k {
	case 4:
		p = 0.05
	case 7:
		p = 0.01
	case 10:
		p = 0.001
	default:
		k = 3
	}
	return p, float64(k), float64(n)
}

func New(k, n, m uint32) (BloomFilter, error) {
	if m == 0 {
		pv, kv, nv := genProbability(k, n)
		mv := -1.0 * ((kv * nv) / math.Log(1-math.Pow(math.E, (math.Log(pv)/kv))))
		k = uint32(kv)
		m = ((uint32(mv) + M_MASK) >> M_OFF) << 10 // * 1024
	}
	return &bloomFilter{
		m:    m,
		k:    k,
		data: make([]byte, m),
	}, nil
}

func (b *bloomFilter) Show() []byte {
	return b.data
}

func (b *bloomFilter) Read(data []byte) ([]byte, error) {
	b.data = data[:b.m]
	return data[b.m:], nil
}

func (b *bloomFilter) Elem(data []byte) bool {
	for i := uint32(0); i < b.k; i++ {
		h := MurmurHash.New(i)
		h.Write(data)
		hv, _ := miscellaneous.D32func(h.Sum(nil))
		if b.getbit(hv) {
			return true
		}
	}
	return false
}

func (b *bloomFilter) Insert(data []byte) error {
	for i := uint32(0); i < b.k; i++ {
		h := MurmurHash.New(i)
		h.Write(data)
		hv, _ := miscellaneous.D32func(h.Sum(nil))
		b.setbit(hv)
	}
	return nil
}
