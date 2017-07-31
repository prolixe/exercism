package paasio

import (
	"io"
	"sync"
)

const testVersion = 3

type Reader interface {
	Read(p []byte) (int, error)
}

type Writer interface {
	Write(p []byte) (int, error)
}
type ReaderWriter interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
}

type Paasio struct {
	buf  ReaderWriter
	m    sync.Mutex
	rc   int64
	wc   int64
	rops int
	wops int
}
type PaasioReader struct {
	buf  Reader
	m    sync.Mutex
	rc   int64
	rops int
}
type PaasioWriter struct {
	buf  Writer
	m    sync.Mutex
	wc   int64
	wops int
}

func NewReadWriteCounter(rw ReaderWriter) ReadWriteCounter {
	return &Paasio{buf: rw}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &PaasioReader{buf: r}
}
func NewWriteCounter(w io.Writer) WriteCounter {
	return &PaasioWriter{buf: w}
}

func (pa *PaasioReader) Read(p []byte) (n int, err error) {
	pa.m.Lock()
	defer pa.m.Unlock()
	n, err = pa.buf.Read(p)
	pa.rops++
	pa.rc += int64(n)
	return
}

func (pa *PaasioWriter) Write(p []byte) (n int, err error) {
	pa.m.Lock()
	defer pa.m.Unlock()
	n, err = pa.buf.Write(p)
	pa.wops++
	pa.wc += int64(n)
	return
}

func (pa *Paasio) Read(p []byte) (n int, err error) {
	pa.m.Lock()
	defer pa.m.Unlock()
	n, err = pa.buf.Read(p)
	pa.rops++
	pa.rc += int64(n)
	return
}

func (pa *Paasio) Write(p []byte) (n int, err error) {
	pa.m.Lock()
	defer pa.m.Unlock()
	n, err = pa.buf.Write(p)
	pa.wops++
	pa.wc += int64(n)
	return
}

func (pa *PaasioReader) ReadCount() (n int64, nops int) {
	return pa.rc, pa.rops
}

func (pa *Paasio) ReadCount() (n int64, nops int) {
	return pa.rc, pa.rops
}

func (pa *PaasioWriter) WriteCount() (n int64, nops int) {
	return pa.wc, pa.wops
}
func (pa *Paasio) WriteCount() (n int64, nops int) {
	return pa.wc, pa.wops
}
