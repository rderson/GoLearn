// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Reverb1 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// !+
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	timeout := time.NewTimer(10 * time.Second)
	text := make(chan string)

	go func ()  {
		for input.Scan() {
			text <- input.Text()
		}
	}()

	for {
		select {
		case t, ok := <- text :
			if ok {
				timeout.Reset(10 * time.Second)
				echo(c, t, 1*time.Second)
			} else {
				c.Close()
				return
			}
		case <-timeout.C :
			timeout.Stop()
			c.Close()
			fmt.Println("Timeout expired. Connection closed.")
			return
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
