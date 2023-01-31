package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferencia é:", area)
}
