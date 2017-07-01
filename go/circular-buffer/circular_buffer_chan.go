package circular

import "errors"

const testVersion = 4

type Buffer struct {
	buf chan byte
	cap int
}

func NewBuffer(size int) *Buffer {
	buf := Buffer{make(chan byte, size), size}
	return &buf
}

func (b *Buffer) ReadByte() (result byte, err error) {
	if len(b.buf) == 0 {
		return result, errors.New("Empty!")
	}
	return <-b.buf, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.buf) == cap(b.buf) {
		return errors.New("Full!")
	}
	b.buf <- c
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		<-b.buf
		b.buf <- c
	}
}
func (b *Buffer) Reset() {
	b.buf = make(chan byte, b.cap)
}
