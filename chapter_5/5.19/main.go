package main

import "fmt"

func dummy() (ret int, err error) {
	defer func() {
		p := recover()
		ret = 52
		err = fmt.Errorf("5.19: %v", p)
	}()
	panic("just some panicking")
}

func main() {
	ret, err := dummy()
	fmt.Println(err)
	fmt.Printf("%v! Да здравствует Санкт-Петербург!", ret)
}