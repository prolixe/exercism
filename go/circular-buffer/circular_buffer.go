package circular

import "errors"

const testVersion = 4

type Buffer struct {
	array []byte
	wPos  int
	rPos  int
	n     int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{array: make([]byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {

	if b.n == 0 {
		return 0, errors.New("Buffer empty")
	}
	c := b.array[b.rPos]
	b.moveReadPointer()
	b.n--
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {

	if b.n == len(b.array) {
		return errors.New("Buffer full")
	}
	b.array[b.wPos] = c
	b.moveWritePointer()
	b.n++
	return nil
}

func (b *Buffer) Overwrite(c byte) {

	if b.n == len(b.array) {
		b.array[b.wPos] = c
		b.moveWritePointer()
		//We must also increase the read position if the buffer is full
		b.moveReadPointer()
	} else {
		b.WriteByte(c)
	}
}

func (b *Buffer) Reset() {
	b.array = make([]byte, len(b.array))
	b.wPos = 0
	b.rPos = 0
	b.n = 0
}

func (b *Buffer) moveReadPointer() {
	b.rPos++
	b.rPos %= len(b.array)
}
func (b *Buffer) moveWritePointer() {
	b.wPos++
	b.wPos %= len(b.array)
}
