package main
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	output, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchAll: %v\n", err)
		os.Exit(-1)
	}
	defer output.Close()
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
		}
		for range os.Args[1:] {
		input := fmt.Sprintln(<-ch) // receive from channel ch
		io.WriteString(output, input)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	io.WriteString(output, "\n")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}