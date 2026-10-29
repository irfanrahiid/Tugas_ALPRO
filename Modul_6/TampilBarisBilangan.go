package main
import "fmt"

func main() {
	var i, j int

	for i = 1; i <= 3; i++ {
		for j = 1; j <= 4; j++ {
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}