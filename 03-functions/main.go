package main

import "fmt"

func sum(a int, b int) int {
	c := 0
	c = a + b
	return c
}

func main() {
	fmt.Println(sum(3, 4))
}
