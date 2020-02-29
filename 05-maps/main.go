package main

import "fmt"

func main() {
	//Define Map
	// emails := make(map[string]string)
	// //Assign value to map
	// emails["Sona"] = "sona@gmail.com"
	// emails["Atul"] = "atul@gmail.com"
	// emails["Shivesh"] = "shivesh@gmail.com"

	//Declare map and add key-value
	emails := map[string]string{"Atul": "atul@gmail.com", "z": "az@gmail.com", "m": "m@gmail.com", "Sona": "sona@gmail.com", "Shivesh": "shivesh@gmail.com"}

	emails["Chris"] = "chris@gmail.com"
	fmt.Println(emails)
	fmt.Println("Length of map:", len(emails))
	fmt.Println(emails["Atul"])

	//Delete from map
	delete(emails, "Shivesh")
	fmt.Println(emails)
}
