package main

import "fmt"

func main() {
	var f int

	fmt.Print("masukan nilai fahrenheit = ")
	fmt.Scan(&f)

	k := (float64(f) + 459.67) * (5.0 / 9.0)
	c := (float64(f) - 32) * (5.0 / 9.0)
	r := c * (4.0 / 5.0)
	
	fmt.Println(k)
	fmt.Println(c)
	fmt.Println(r)
}