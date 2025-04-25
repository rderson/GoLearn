package main

import "fmt"

func counter(out chan<- int)  {
	for x:=0; x <= 10; x++ {
		out <- x
	}
	close(out)
}

func stage(in <-chan int, out chan<-int, fn func(int) int)  {
	for x := range in {
		out <- fn(x)
	}
	close(out)
}

func printer(in <-chan int)  {
	for i := range in {
		fmt.Println("Result: ", i)
	}
}

func main() {
	n := 9000000

	fns := make([]func(int)int, n)
	for i := 0; i < n; i++ {
		fns[i] = func(i int) int {return i+1}
	}

	channels := make([]chan int, n+1)
	for i := range channels {
		channels[i] = make(chan int)
	}

	go counter(channels[0])
	for i := 0; i < n; i++ {
		go stage(channels[i], channels[i+1], fns[i])
	}
	printer(channels[len(channels)-1])
}