package MurmurHash

const (
	m  uint32 = 5
	n  uint32 = 0xe6546b64
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
)

type digest struct {
	h      uint32
	seed   uint32
	length uint64
	tail   []byte
}
