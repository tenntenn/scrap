package a

import (
	"fmt"
	_ "io"
	"reflect"
	"testing"
	_ "unsafe"
)

//go:linkname MyLimitedReader io._LimitedReader
type MyLimitedReader interface {
	R Reader // underlying reader
	N int64  // max bytes remaining
}

func (l *MyLimitedReader) Read(p []byte) (n int, err error) {
	panic("a")
}

func Test(t *testing.T) {
	r := io.LimitReader{
	var w MyWriter
	fmt.Printf("%v\n", reflect.TypeOf(w))
}
