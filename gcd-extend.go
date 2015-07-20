package main

import (
	"fmt"
)

func gcd_extend(a, b int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	} else {
		x, y, q := gcd_extend(b, a%b)
		x, y = y, (x - (a/b)*y)
		return x, y, q
	}
}

func main() {
	x, y, q := gcd_extend(8, 3)
	fmt.Printf("8, 3 扩展欧几里的值: x=%v, y=%v, q=%v\n", x, y, q)
	x, y, q = gcd_extend(11, 6)
	fmt.Printf("11, 6 扩展欧几里的值: x=%v, y=%v, q=%v\n", x, y, q)
	x, y, q = gcd_extend(22, 33)
	fmt.Printf("22, 33 扩展欧几里的值: x=%v, y=%v, q=%v\n", x, y, q)
}
