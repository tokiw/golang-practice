package bzip

import "C"

import (
	"io"
	"sync"
	"unsafe"
)

type writer struct {
	w      io.Writer // underlying output stream
	stream *C.bz_stream
	outbuf [64 * 1024]byte
	m      sync.Mutex
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	const blockSize = 9
	const verbosity = 0
	const workFactor = 30
	w := &writer{w: out, stream: C.bz2alloc()}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

//!-

//!+write
func (w *writer) Write(data []byte) (int, error) {
	w.m.Lock()
	defer w.m.Unlock()
	if w.stream == nil {
		panic("closed")
	}
	var total int // uncompressed bytes written

	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(inlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}

//!-write

//!+close
// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.m.Lock()
	defer w.m.Unlock()
	if w.stream == nil {
		panic("closed")
	}
	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		C.bz2free(w.stream)
		w.stream = nil
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}
