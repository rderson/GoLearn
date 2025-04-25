package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
	City string
}

var people = []*Person{
	{"Maria", 27, "Luebeck"},
	{"Islam", 25, "Mahachkala"},
	{"Rasul", 16, "Nasran"},
	{"Artur", 52, "Moscow"},
	{"Denis", 29, "Saint Petersburg"},
	{"Anna", 22, "Luebeck"},
	{"Oleg", 35, "Rostov"},
	{"Vera", 30, "Ufa"},
	{"Anna", 30, "Berlin"},
	{"Maria", 30, "Berlin"},
	{"Denis", 25, "Moscow"},
	{"Artur", 17, "Moscow"},
	{"Vera", 22, "Saint Petersburg"},
	{"Rasul", 52, "Mahachkala"},
}

type multTier struct {
	p 			[]*Person
	primary		string
	secondary	string
	third		string
}

func (mt *multTier) Len() int					{ return len(mt.p) }
func (mt *multTier) Swap(x, y int)				{ mt.p[x], mt.p[y] = mt.p[y], mt.p[x]}
func (mt *multTier) Less(x, y int) bool			{ 
	key := mt.primary
	for k := 0; k < 3; k++ {
		switch key {
		case "Name":
			if mt.p[x].Name !=  mt.p[y].Name {
				return mt.p[x].Name <  mt.p[y].Name
			}
		case "Age":
			if mt.p[x].Age !=  mt.p[y].Age {
				return mt.p[x].Age <  mt.p[y].Age
			}
		case "City":
			if mt.p[x].City !=  mt.p[y].City {
				return mt.p[x].City <  mt.p[y].City
			}
		}
		if k == 0 {
			key = mt.secondary
		} else if k == 1 {
			key = mt.third
		}
	}
	return false
}

func setPrimary(x *multTier, p string)  {
	x.primary, x.secondary, x.third = p, x.primary, x.secondary
}

func SetPrimary(x sort.Interface, p string)  {
	switch x := x.(type) {
	case *multTier:
		setPrimary(x, p)
	}
}

func NewMT(p []*Person, pr, s, t string) sort.Interface {
	return &multTier{
		p: p,
		primary: pr,
		secondary: s,
		third: t,
	}
}

func main()  {
	fmt.Println("\nMultier:")
	multi := NewMT(people, "Name", "", "")
	sort.Sort(multi)
	for _, person := range people {
		fmt.Printf("%s\t%d\t%s\n", person.Name, person.Age, person.City)
	}

	fmt.Println()
	SetPrimary(multi, "Age")
	sort.Sort(multi)
	for _, person := range people {
		fmt.Printf("%s\t%d\t%s\n", person.Name, person.Age, person.City)
	}

	fmt.Println()
	SetPrimary(multi, "City")
	sort.Sort(multi)
	for _, person := range people {
		fmt.Printf("%s\t%d\t%s\n", person.Name, person.Age, person.City)
	}
}