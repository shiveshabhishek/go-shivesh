package main

import (
	"fmt"
	"github.com/shiveshabhishek/go-shivesh/04-arrays_slices/pkg"
)

func main() {
	arr := array.CreateArrayWithEntries()
	fmt.Println("Array Entries:", arr)
	minValue,er := array.MinVal(arr)
	fmt.Println("Minimum Value:", minValue,er)
	maxValue,er := array.MaxVal(arr)
	fmt.Println("Maximum Value:", maxValue,er)
}
