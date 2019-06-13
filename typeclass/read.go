package typeclass

// read失败时返回错误信息，出错时返回[]byte{}和出错信息
type Read interface {
	Read([]byte) ([]byte, error) // 返回尚未解析的字节数组
}
