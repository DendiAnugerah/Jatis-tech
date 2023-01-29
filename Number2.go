package main

import "fmt"

func NomerDua(input string) string {
	result := ""
	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}
	return result
}

func main() {
	fmt.Println(NomerDua("jatis"))
}
