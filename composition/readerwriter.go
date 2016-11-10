package main

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}
type Buffer struct {
	rw          ReadWriter
	numPosition int
}

func (b *Buffer) ReadBuffer() error {
	for n := 0; n < numPostion; n++ {
		b.rw.Read()
	}
}
func (b *Buffer) WriteBuffer() error {
	for n := 0; n < numPostion; n++ {
		b.rw.Write()
	}
}
