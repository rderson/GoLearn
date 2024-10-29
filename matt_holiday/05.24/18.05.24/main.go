package  main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name string
	Number int
	Boss *Employee
	Hired time.Time
}

func main() {
	var e Employee

	e.Name = "4ort"
	e.Number = 1
	e.Hired = time.Now()

	var b = Employee{
		"Bot",
		2,
		&e,
		time.Now().AddDate(1, 6, 26),
	}

	c := Employee{
		Name: "Botik",
		Number: 3,
		Boss: &b,
		Hired: time.Now().AddDate(1, 6, 27),
	}

	fmt.Printf("%T %+[1]v\n", e)
	fmt.Printf("%T %+[1]v\n", b)
	fmt.Printf("%T %+[1]v\n",     c)

	d := map[string]*Employee{}	

	d["Spear Shot"] = &Employee{"Spear Shot", 4, nil, time.Now()}

	d["Brohan"] = &Employee{
		Name: "Brohan",
		Number: 5,
		Boss: d["Spear Shot"],
		Hired: time.Now().AddDate(0, 0, 1),
	}

	fmt.Printf("%T %+[1]v\n", d)

}