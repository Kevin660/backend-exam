package main

import (
	"fmt"
)

func swap[T any](a, b *T) {
	if a == nil || b == nil {
		panic("a or b is nil")
	}
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)
	// swap(&a, nil)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
