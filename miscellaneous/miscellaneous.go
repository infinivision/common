package miscellaneous

import (
	"errors"

	"github.com/infinivision/common/slice"
	"github.com/infinivision/common/typeclass"
)

type uint8ToByteFunc (func(uint8) []byte)
type byteToUint8Func (func([]byte) (uint8, error))

type uint16ToByteFunc (func(uint16) []byte)
type byteToUint16Func (func([]byte) (uint16, error))

type uint32ToByteFunc (func(uint32) []byte)
type byteToUint32Func (func([]byte) (uint32, error))

type uint64ToByteFunc (func(uint64) []byte)
type byteToUint64Func (func([]byte) (uint64, error))

var E8func uint8ToByteFunc
var D8func byteToUint8Func
var E16func uint16ToByteFunc
var D16func byteToUint16Func
var E32func uint32ToByteFunc
var D32func byteToUint32Func
var E64func uint64ToByteFunc
var D64func byteToUint64Func

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

func init() { // 初始化整型序列化函数
	buf := uint32ToBytes(1)
	E8func = uint8ToBytes
	D8func = bytesToUint8
	if buf[0] == 0 {
		E16func = bigEndianUint16ToBytes
		D16func = bigEndianBytesToUint16
		E32func = bigEndianUint32ToBytes
		D32func = bigEndianBytesToUint32
		E64func = bigEndianUint64ToBytes
		D64func = bigEndianBytesToUint64
	} else {
		E16func = uint16ToBytes
		D16func = bytesToUint16
		E32func = uint32ToBytes
		D32func = bytesToUint32
		E64func = uint64ToBytes
		D64func = bytesToUint64
	}
}

// big endian to little endian
func bigEndianUint16ToBytes(a uint16) []byte {
	buf := make([]byte, 2)
	buf[1] = byte(a & 0xFF)
	buf[0] = byte((a >> 8) & 0xFF)
	return buf
}

func bigEndianBytesToUint16(a []byte) (uint16, error) {
	if len(a) != 2 {
		return 0, errors.New("bigEndianBytesToUint16: Illegal slice length")
	}
	b := uint16(0)
	for i, v := range a {
		b += uint16(v)
		if i > 0 {
			b <<= 8
		}
	}
	return b, nil
}

// big endian to little endian
func bigEndianUint32ToBytes(a uint32) []byte {
	buf := make([]byte, 4)
	buf[3] = byte(a & 0xFF)
	buf[2] = byte((a >> 8) & 0xFF)
	buf[1] = byte((a >> 16) & 0xFF)
	buf[0] = byte((a >> 24) & 0xFF)
	return buf
}

func bigEndianBytesToUint32(a []byte) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("bigEndianBytesToUint32: Illegal slice length")
	}
	b := uint32(0)
	for i, v := range a {
		b += uint32(v)
		if i > 0 {
			b <<= 8
		}
	}
	return b, nil
}

// big endian to little endian
func bigEndianUint64ToBytes(a uint64) []byte {
	buf := make([]byte, 8)
	buf[7] = byte(a & 0xFF)
	buf[6] = byte((a >> 8) & 0xFF)
	buf[5] = byte((a >> 16) & 0xFF)
	buf[4] = byte((a >> 24) & 0xFF)
	buf[3] = byte((a >> 32) & 0xFF)
	buf[2] = byte((a >> 40) & 0xFF)
	buf[1] = byte((a >> 48) & 0xFF)
	buf[0] = byte((a >> 56) & 0xFF)
	return buf
}

func bigEndianBytesToUint64(a []byte) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("bigEndianBytesToUint64: Illegal slice length")
	}
	b := uint64(0)
	for i, v := range a {
		b += uint64(v)
		if i > 0 {
			b <<= 8
		}
	}
	return b, nil
}

func uint8ToBytes(a uint8) []byte {
	buf := make([]byte, 1)
	buf[0] = a & 0xFF
	return buf
}

func bytesToUint8(a []byte) (uint8, error) {
	if len(a) != 1 {
		return 0, errors.New("BytesToUint8: Illegal slice length")
	}
	return uint8(a[0]), nil
}

func uint16ToBytes(a uint16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(a & 0xFF)
	buf[1] = byte((a >> 8) & 0xFF)
	return buf
}

func bytesToUint16(a []byte) (uint16, error) {
	if len(a) != 2 {
		return 0, errors.New("BytesToUint16: Illegal slice length")
	}
	b := uint16(0)
	for i, v := range a {
		b += uint16(v) << (8 * uint16(i))
	}
	return b, nil
}

func uint32ToBytes(a uint32) []byte {
	buf := make([]byte, 4)
	buf[0] = byte(a & 0xFF)
	buf[1] = byte((a >> 8) & 0xFF)
	buf[2] = byte((a >> 16) & 0xFF)
	buf[3] = byte((a >> 24) & 0xFF)
	return buf
}

func bytesToUint32(a []byte) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("bytesToUint32: Illegal slice length")
	}
	b := uint32(0)
	for i, v := range a {
		b += uint32(v) << (8 * uint32(i))
	}
	return b, nil
}

func uint64ToBytes(a uint64) []byte {
	buf := make([]byte, 8)
	buf[0] = byte(a & 0xFF)
	buf[1] = byte((a >> 8) & 0xFF)
	buf[2] = byte((a >> 16) & 0xFF)
	buf[3] = byte((a >> 24) & 0xFF)
	buf[4] = byte((a >> 32) & 0xFF)
	buf[5] = byte((a >> 40) & 0xFF)
	buf[6] = byte((a >> 48) & 0xFF)
	buf[7] = byte((a >> 56) & 0xFF)
	return buf
}

func bytesToUint64(a []byte) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("bytesToUint64: Illegal slice length")
	}
	b := uint64(0)
	for i, v := range a {
		b += uint64(v) << (8 * uint64(i))
	}
	return b, nil
}

func Dup(a []byte) []byte {
	if a == nil {
		return nil
	}
	b := []byte{}
	return append(b, a...)
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
