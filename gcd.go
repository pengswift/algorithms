package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	fmt.Printf("6, 8 的最大公约数: %v\n", gcd(6, 8))
	fmt.Printf("13, 5 的最大公约数: %v\n", gcd(13, 5))
}
