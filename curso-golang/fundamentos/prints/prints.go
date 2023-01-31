package main

import "fmt"

func main() {
	x := 3.145
	xs := fmt.Sprint(x)
	fmt.Println("o valor de x é: " + xs + ".")
	fmt.Println("o valor de x é", xs, ".")
	//fmt.Printf("o valor de x %.f", x)
	fmt.Printf("o valor de x %.2f.", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
