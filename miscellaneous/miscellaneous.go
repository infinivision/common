package miscellaneous

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/infinivision/common/slice"
	"github.com/infinivision/common/typeclass"
)

func Max(x, y int) int {
	switch {
	case x < y:
		return y
	default:
		return x
	}
}

func Min(x, y int) int {
	switch {
	case x < y:
		return x
	default:
		return y
	}
}

func Dup(a []byte) []byte {
	if a == nil {
		return nil
	}
	b := []byte{}
	return append(b, a...)
}

func E8func(a uint8) []byte {
	buf := make([]byte, 1)
	buf[0] = a & 0xFF
	return buf
}

func D8func(a []byte) (uint8, error) {
	if len(a) != 1 {
		return 0, errors.New("D8func: Illegal slice length")
	}
	return uint8(a[0]), nil
}

func E16func(a uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, a)
	return buf
}

func D16func(a []byte) (uint16, error) {
	if len(a) != 2 {
		return 0, errors.New("D16func: Illegal slice length")
	}
	return binary.LittleEndian.Uint16(a), nil
}

func E32func(a uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, a)
	return buf
}

func D32func(a []byte) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("D32func: Illegal slice length")
	}
	return binary.LittleEndian.Uint32(a), nil
}

func E64func(a uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, a)
	return buf
}

func D64func(a []byte) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("D64func: Illegal slice length")
	}
	return binary.LittleEndian.Uint64(a), nil
}

func EB16func(a uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, a)
	return buf
}

func DB16func(a []byte) (uint16, error) {
	if len(a) != 2 {
		return 0, errors.New("DB16func: Illegal slice length")
	}
	return binary.BigEndian.Uint16(a), nil
}

func EB32func(a uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, a)
	return buf
}

func DB32func(a []byte) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("DB32func: Illegal slice length")
	}
	return binary.BigEndian.Uint32(a), nil
}

func EB64func(a uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, a)
	return buf
}

func DB64func(a []byte) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("DB64func: Illegal slice length")
	}
	return binary.BigEndian.Uint64(a), nil
}

// slice length || slice
func Eslice(a []byte) []byte {
	buf := []byte{}
	buf = append(buf, E32func(uint32(len(a)))...)
	buf = append(buf, a...)
	return buf
}

func Dslice(data []byte) ([]byte, []byte, error) {
	if len(data) < 4 {
		return []byte{}, []byte{}, errors.New("DecodeSlice: Illegal slice length")
	}
	n, _ := D32func(data[:4])
	if data = data[4:]; uint32(len(data)) < n {
		return []byte{}, []byte{}, errors.New("DecodeSlice: Illegal slice length")
	}
	return Dup(data[:n]), data[n:], nil
}

// xs是否包含ys, xs和ys均为排序后的数组
func Contain(xs, ys []typeclass.Ord) bool {
	switch {
	case len(ys) == 0:
		return true
	case len(xs) < len(ys):
		return false
	case xs[len(ys)-1].Gt(ys[len(ys)-1]):
		return false
	case len(ys) == 1:
		if slice.Bsearch(ys[0], xs) != -1 {
			return true
		}
		return false
	case len(ys) == 2:
		if slice.Bsearch(ys[0], xs) != -1 && slice.Bsearch(ys[1], xs) != -1 {
			return true
		}
		return false
	default:
		start := slice.Bsearch(ys[0], xs)
		if start == -1 {
			return false
		}
		end := slice.Bsearch(ys[len(ys)-1], xs)
		if end == -1 {
			return false
		}
		xs = xs[start:end]
		for i, j := 1, len(ys)-1; i < j; i++ {
			start = slice.Bsearch(ys[i], xs)
			if start == -1 {
				return false
			}
			xs = xs[start:]
		}
		return true
	}
}

func MostFrequent(xs []typeclass.Ord) (int, []typeclass.Ord) {
	switch {
	case len(xs) == 0:
		return 0, []typeclass.Ord{}
	case len(xs) == 1:
		return 1, []typeclass.Ord{xs[0]}
	}
	n, m := 1, 0
	rs := []typeclass.Ord{}
	for i, j := 0, len(xs); i < j; i++ {
		if i == j-1 || xs[i].NotEq(xs[i+1]) {
			switch {
			case n == m:
				rs = append(rs, xs[i])
			case n > m:
				m = n
				rs = []typeclass.Ord{xs[i]}
			}
			n = 0
		}
		n++
	}
	return m, rs
}

func Encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}
