package gee_cache

// ByteView 用来标识缓存值，只读
type ByteView struct {
	b []byte
}

// Len 计算对象的字节数
func (byteView ByteView) Len() int64 {

	return 0
}

// ByteSlice 返回一个拷贝值
func (byteView ByteView) ByteSlice() []byte {

	return nil
}

// String 返回字符串表示
func (byteView ByteView) String() string {

	return ""
}

func (byteView ByteView) cloneBytes(b []byte) []byte {

	return nil
}