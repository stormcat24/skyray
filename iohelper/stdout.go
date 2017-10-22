package iohelper

import "io"

type BufferObserver struct {
	callback func(d []byte)
}

func NewBufferObserver(callback func(d []byte)) io.Writer {
	return &BufferObserver{
		callback: callback,
	}
}

func (b *BufferObserver) Write(p []byte) (n int, err error) {
	b.callback(p)
	return len(p), nil
}
