package main

import "fmt"

func main() {
	a := 3
	b := 2
	//bitwise
	fmt.Println("e =>", a&b)   // 11 & 10 = 10
	fmt.Println("ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("xor =>", a^b) // 11 ^ 10 = 01
}
