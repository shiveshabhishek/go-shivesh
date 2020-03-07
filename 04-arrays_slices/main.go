package array

import "fmt"

func main() {
	// Arrays can also be declared like below:
	// var vehicleArr [2]string
	// vehicleArr[0] = "BMW"
	// vehicleArr[1] = "Ford"

	// Short-hand for arrays in go:
	vehicleArr := [2]string{"BMW", "FORD"}                   // This is an Array
	fruitArr := []string{"Apple", "Grape", "Banana", "Kiwi"} // This is a Slice
	fmt.Println(vehicleArr)
	// vehicleArr=append(vehicleArr,"LAMBO")
	//--> IF USED APPEND IN ARRAY AS ABOVE, IT PRODUCES FOLLOWING ERROR:
	//--->  ./main.go:15:19: first argument to append must be slice; have [2]string
	fmt.Println("Length of fruitArr before appending:", len(fruitArr))
	fruitArr = append(fruitArr, "Litchi")
	fmt.Println(fruitArr)
	fmt.Println("Length of fruitArr after appending Litchi: ", len(fruitArr))
}
