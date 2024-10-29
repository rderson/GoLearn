package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriter struct {
	writer io.Writer
	count int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	cw.count += int64(n)
	return n, err
} 

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{writer: w}
	return cw, &cw.count
}

func main() {
	var buf io.Writer = &bytes.Buffer{}
	writer, count := CountingWriter(buf)

	writer.Write([]byte("AHAHAHAHHAHAHAH"))
	writer.Write([]byte("Ti ochkochnik"))

	fmt.Println("Bytes written: ", *count)
}