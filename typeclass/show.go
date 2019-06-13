package typeclass

// show失败时返回[]byte{}
type Show interface {
	Show() []byte
}
