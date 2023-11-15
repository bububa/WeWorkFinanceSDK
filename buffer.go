package wxworkfinancesdk

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func BufferPool() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func BufferPoolRelease(v *bytes.Buffer) {
	v.Reset()
	bufferPool.Put(v)
}

var readerPool = sync.Pool{
	New: func() any {
		return new(bytes.Reader)
	},
}

func ReaderPool(bs []byte) *bytes.Reader {
	r := readerPool.Get().(*bytes.Reader)
	r.Reset(bs)
	return r
}

func ReaderPoolRelease(r *bytes.Reader) {
	readerPool.Put(r)
}
