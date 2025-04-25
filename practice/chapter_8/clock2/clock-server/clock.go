// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, tz *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(tz).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	zone := flag.String("zone", "Local", "Name of the time zone (e.g., Asia/Tokyo)")
	port := flag.String("port", "8000", "Port to listen on")
	flag.Parse()


	loc, err := time.LoadLocation(*zone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "clock2: %v", err)
		return
	}

	addr := "localhost:" + *port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, loc) // handle connections concurrently
	}
	//!-
}
