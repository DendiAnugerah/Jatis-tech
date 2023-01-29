package main

import "fmt"

func NomerSatu(a *int, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}

func main() {
	a := 50
	b := 63
	fmt.Println(NomerSatu(&a, &b))
}
