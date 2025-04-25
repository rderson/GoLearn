package main

import (
	"fmt"
	"sync"
	"time"
)


func main()  {
	var wg sync.WaitGroup

	wg.Add(2)

	const duration = 1*time.Second
	var count int
	pingCh := make(chan struct{})
	pongCh := make(chan struct{})
	
	fmt.Println("The game starts!")
	start := time.Now()

	go func ()  {
		defer wg.Done()
		for {
			_, ok := <-pingCh
			if !ok {
				return
			}
			fmt.Println("Ping...")
			select {
			case pongCh <- struct{}{}:
			default:
				return
			}
		}
	}()
	go func ()  {
		defer wg.Done()
		for time.Since(start) < duration {
			<-pongCh
			fmt.Println("Pong!")
			pingCh <- struct{}{}
			count += 2
		}
		close(pingCh)
	}()

	pingCh <- struct{}{}
	wg.Wait()
	fmt.Println("The game is finished!")
	fmt.Println("Total number of ping-pongs during one second: ", count)
	
}

