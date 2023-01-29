package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NomerTiga(input string) string {
	var output string
	seen := make(map[rune]bool)

	for _, char := range input {
		if !seen[char] {
			seen[char] = true
			count := 0
			for _, nextChar := range input[1:] {
				if char == nextChar {
					count++
				}
			}
			if count > 1 {
				output += strconv.Itoa(count)
			}
			output += string(strings.ToLower(string(char)))
		}
	}

	return strings.ReplaceAll(output, " ", "")
}

func main() {
	fmt.Println(NomerTiga("dani Maulana"))
}
