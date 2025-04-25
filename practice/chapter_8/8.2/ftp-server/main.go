package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func handleConnection(c net.Conn)  {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		cmds := strings.Split(scanner.Text(), " ")
		command := cmds[0]

		switch command {
		case "cd":
			if len(cmds) != 2 {
				log.Fatal("Usage: cd directory")
			}
			if err := os.Chdir(cmds[1]); err != nil {
				log.Print(err)
			}
		case "ls":
			exeCmd(command, cmds[1:]...)
		case "get":
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Print(err)
			}
			data, err := io.ReadAll(file)
			if err != nil {
				log.Print(err)
			}
			fmt.Fprintf(c, "%s\n", string(data)) 
		case "close":
			log.Print("Connection closed.")
			return
		default:
			log.Print("Unknown command.")
			fmt.Fprint(c, "ls: list content\ncd: change directory\nget: get content of the selected file\nclose: close the connection\n")
		}
	}
}

func exeCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}