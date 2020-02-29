package main

import "fmt"

func main() {
	// Arrays can also be declared like below:
	// var vehicleArr [2]string
	// vehicleArr[0] = "BMW"
	// vehicleArr[1] = "Ford"

	// Short-hand for arrays in go:
	vehicleArr := [2]string{"BMW", "FORD"}
	fruitArr := []string{"Apple", "Grape", "Banana", "Kiwi"}
	fmt.Println(vehicleArr)
	fmt.Println(fruitArr)
}
