//program ini saya buat untuk menghitung fungsi f(x,y)

package main

import "fmt"

func main() {
	var f, x, y float64

	fmt.Print("masukan nilai x = ")
	fmt.Scanln(&x)

	fmt.Print("masukan nilai y = ")
	fmt.Scanln(&y)

	f = (1 / (3 * x * x + 10)) + (10 * y) + 7

	fmt.Print("f = ", f)
}