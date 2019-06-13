/*
 * algorithmic refer: https://en.wikipedia.org/wiki/MurmurHash
 */
package MurmurHash

import (
	"hash"

	"github.com/infinivision/common/miscellaneous"
)

func New(seed uint32) hash.Hash32 {
	d := new(digest)
	d.seed = seed
	d.Reset()
	return d
}

func (d *digest) Size() int { return 4 }

func (d *digest) BlockSize() int { return 4 }

func (d *digest) Reset() {
	d.h = d.seed
	d.length = 0
	d.tail = []byte{}
}

func (d *digest) Write(p []byte) (int, error) {
	n := len(p)
	d.length += uint64(n)
	if x := len(d.tail); x != 0 {
		n += x
		p = append(d.tail, p...)
		d.tail = []byte{}
	}
	nblocks := n / d.BlockSize()
	d.block(p, nblocks)
	if n&3 != 0 {
		d.tail = append(d.tail, p[nblocks*d.BlockSize():]...)
	}
	return n, nil
}

func (d *digest) Sum(p []byte) []byte {
	d.Write(p)
	if i := d.length & 3; i != 0 {
		d.tail = append(d.tail, zeroBytesSlice()[:4-i]...)
		k, _ := miscellaneous.D32func(d.tail)
		k *= c1
		k = (k << 15) | (k >> 17)
		k *= c2
		d.h ^= k
	}
	d.h ^= uint32(d.length)
	d.h ^= d.h >> 16
	d.h *= 0x85ebca6b
	d.h ^= d.h >> 13
	d.h *= 0xc2b2ae35
	d.h ^= d.h >> 16
	return miscellaneous.E32func(d.h)
}

func (d *digest) Sum32() uint32 {
	d.Sum(nil)
	return d.h
}

// 4 bytes
func zeroBytesSlice() []byte {
	return []byte{0, 0, 0, 0}
}

func (d *digest) block(p []byte, nblocks int) {
	for i, j := 0, nblocks; i < j; i++ {
		k, _ := miscellaneous.D32func(p[i*4 : i*4+4])
		k *= c1
		k = (k << 15) | (k >> 17)
		k = k * c2
		d.h ^= k
		d.h = (d.h << 13) | (d.h >> 19)
		d.h = d.h*m + n
	}
}
