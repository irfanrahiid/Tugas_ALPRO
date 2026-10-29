package main

import "fmt"

func main() {
	var n, j int

	fmt.Print("masukan nilai N = ")
	fmt.Scan(&n)

	for j = 1; j <= n; j++ {
		fmt.Print(j, " * ", j, " = ")
		fmt.Println(j * j)
	}
}