package gee_cache

import (
	"fmt"
	"reflect"
	"testing"
)

func newByteView(b []byte) ByteView {
	return ByteView{
		b: b,
	}
}
func TestByteView_Len(t *testing.T) {
	v := "mkii"
	vByte := []byte(v)
	byteView := newByteView(vByte)

	if byteView.Len() != int64(len(vByte)) {
		t.Fatal("Failed to get length for ByteView")
	}
}

func TestByteView_ByteSliceSlice(t *testing.T) {
	v := "mkii"
	vByte := []byte(v)
	byteView := newByteView(vByte)

	vByteP := fmt.Sprintf("%p", vByte)
	byteViewP := fmt.Sprintf("%p", byteView.b)

	copyByte := byteView.ByteSlice()
	copyP := fmt.Sprintf("%p", copyByte)

	if vByteP != byteViewP || copyP == byteViewP || !reflect.DeepEqual(vByte, copyByte) {
		t.Fatal("Failed to copy from byteView")
	}

	t.Log(vByteP)
	t.Log(byteViewP)
	t.Log(copyP)
}

func TestByteView_String(t *testing.T) {
	v := "mkii"
	vByte := []byte(v)
	byteView := newByteView(vByte)

	if v != byteView.String() {
		t.Fatal("Failed to get string by byteView")
	}
}
