package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type CityServer struct {
	city   string
	server string
	time   string
}

func connectToServer(cs *CityServer, wg *sync.WaitGroup, times []string, index int) {
	defer wg.Done()

	conn, err := net.Dial("tcp", cs.server)
	if err != nil {
		log.Printf("Error connecting to server %s: %v", cs.server, err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		times[index] = scanner.Text()
	}
}

func main() {
	usage := "Usage: ./clockwall city1=server1 city2=server2 city3=server3..."
	timezones := os.Args[1:]
	var cityServers []*CityServer

	for _, timezone := range timezones {
		if strings.Contains(timezone, "=") {
			tz := strings.SplitN(timezone, "=", 2)
			city := tz[0]
			if strings.Contains(city, "_") {
				city = strings.ReplaceAll(city, "_", " ")
			}
			cityServers = append(cityServers, &CityServer{
				city:   city,
				server: tz[1],
			})
		} else {
			fmt.Println("Wrong input.")
			fmt.Println(usage)
			os.Exit(1)
		}
	}

	for _, cs := range cityServers {
		fmt.Printf("|  %-8s  ", cs.city)
	}
	fmt.Println("|")

	times := make([]string, len(cityServers))
	var wg sync.WaitGroup

	for i, cs := range cityServers {
		wg.Add(1)
		go connectToServer(cs, &wg, times, i)
	}

	for {
		for _, time := range times {
			fmt.Printf("|  %-8s  ", time)
		}
		fmt.Println("|")
		time.Sleep(1 * time.Second)
	}
}