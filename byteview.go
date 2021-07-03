package gee_cache

// ByteView 用来标识缓存值，只读
type ByteView struct {
	b []byte
}

// Len 计算对象的字节数
func (byteView ByteView) Len() int64 {
	return int64(len(byteView.b))
}

// ByteSlice 返回一个拷贝值
func (byteView ByteView) ByteSlice() []byte {
	return byteView.cloneBytes(byteView.b)
}

// String 返回字符串表示
func (byteView ByteView) String() string {
	return string(byteView.b)
}

func (byteView ByteView) cloneBytes(b []byte) []byte {
	cpByte := make([]byte, len(b))
	copy(cpByte, b)
	return cpByte
}
