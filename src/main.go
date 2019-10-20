package main
import "fmt"

func main() {
	var g Graph

	g.CreateRandom(10, "test")
	fmt.Println(g)
}