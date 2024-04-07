package main

import (
	"io"
)

type MyReader struct {
	data []byte
	pos  int
}

type MyWriter struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
		data: buffer,
	}
}

func NewWriterToBuffer(buffer []byte) *MyWriter {
	return &MyWriter{
		data: buffer,
	}
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	if len(r.data) == r.pos {
		return 0, io.EOF
	}

	n = len(p)
	if (len(r.data) - r.pos) < n {
		n = len(r.data) - r.pos
	}

	for i := 0; i < n; i++ {
		p[i] = byte('a' + (r.pos+i)%26)
	}

	r.pos += n

	return
}

func (w *MyWriter) Write(p []byte) (n int, err error) {

	w.data = append(w.data, p...)

	w.pos += len(p)

	return len(p), nil
}

func main() {
	w := NewWriterToBuffer(make([]byte, 1024))
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	_, err := w.Write(data)
	if err != nil {
		return
	}
}
