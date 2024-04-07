package main

import (
	"fmt"
	"io"
)

/**
Напишите программу, которая имплементирует io.Reader. Реализуйте работу с чтением из буфера в памяти. Пример - простое чтение строк и их вывод в консоль.
Когда может быть применимо: сделать мок для io.Reader. Например, для сетевого вызова, т.к. io.Reader - просто интерфейс, а чтение может происходить из файла, из сети, etc.
При реализации метода Read учтите корректность работы с смещениями (смещение относительно начала buffer).
*/

type MyReader struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
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

	fmt.Printf("Data: %s", p)

	return
}

func main() {
	buffer := make([]byte, 1024)
	r := NewReaderFromBuffer(buffer)

	data := []byte("abcdefghijklmnopqrstuvwxyz")
	rd, err := r.Read(data)
	fmt.Printf("Len: %d; Len read: %d", len(buffer), rd)
	if err != nil {
		return
	}
}
