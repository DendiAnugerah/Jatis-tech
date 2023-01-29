package main

import "fmt"

func NomerEmpat(input int) {
	count := 1

	for i := 1; i <= input; i++ {
		if isPrime(i) {
			fmt.Println("Bilangan prima:", i)
		} else {
			if i%9 == 0 {
				fmt.Println("Bilangan kelipatan 9 ke-",count)
				count++
			} else {
				fmt.Println(i)
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	NomerEmpat(100)
}
