package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getOne(i int) []byte{
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main()  {
	var (
		output io.WriteCloser = os.Stdout
		err error
		count int
		fails int
		data []byte
	) 

	if len(os.Args) > 0 {
		output, err = os.Create(os.Args[1])

		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; fails < 2; i++{
		if data = getOne(i); data == nil{
			fails++
			continue
		}

		if count > 0{
			fmt.Fprint(output, ",")
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed: %s\n", err)
			os.Exit(-1)
		}

		fails = 0
		count++
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", count)
}