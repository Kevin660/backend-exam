package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name string
	Age  int
	Next *Person
}

func TrimAllStrings(a *Person) {
	p := a
	visited := make(map[uintptr]bool)
	for {
		if p == nil {
			break
		}

		ptr := reflect.ValueOf(p).Pointer()
		if visited[ptr] {
			break
		}

		visited[ptr] = true
		p.Name = strings.TrimSpace(p.Name)
		p = p.Next
	}
}

func main() {
	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(a)
	fmt.Println(a.Next.Next.Name == "name")
}
