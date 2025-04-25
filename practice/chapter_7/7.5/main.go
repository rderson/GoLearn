// Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r and a number of bytes n,
// and returns another Reader that reads from r but reports an end-of-file condition after n bytes.
// Implement it.

package main

import (
	"fmt"
	"io"
	"strings"
)

// creating LimitedReader struct that we can return with our implemented function
type LimitedReader struct {
	r io.Reader 	// reader from which we read
	n int64			// remaining amount of bytes 
}

// implementing LimitReader function
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

// implementing Read method to satisfy io.Reader interface
func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	// if n is lower or equal to zero return EOF (n - remaining amount of bytes)
	if lr.n <= 0 {
		return 0, io.EOF
	}

	// limit the slice p to n bytes, ensuring we do not read more than the remaining amount of bytes
	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}

	n, err = lr.r.Read(p)	// reading from the original reader that is lr.r into limited byte slice p
	lr.n -= int64(n)		// updating number lr.n by subtracting bytes we have already read
	return
}

func main() {
	data := "Just an example." 							// define a string that stores desiered data
	reader := LimitReader(strings.NewReader(data), 4)	// create a limited reader that reads 4 bytes from our data string

	buf := make([]byte, 8)								// create a buffer with a capacity larger than the limit
	n, err := reader.Read(buf)							// read from the limited reader
	fmt.Printf("Read %d bytes: %q (error: %v)\n", n, buf[:n], err) // print the bytes read and an error (expected: nil or io.EOF)
}