package main

import "fmt"

func main() {
	nameArray := []string{"Pradeep", "Shivesh", "Shubham", "Sona"}
	fmt.Printf("Names:\n")
	for i, name := range nameArray {
		i++
		fmt.Printf("%d %s\n", i, name)
	}
	bdays := map[string]string{"Shivesh": "11-Oct", "Pradeep": "28-Jan", "Shubham": "3 Mar", "Sona": "5-Sep"}
	fmt.Println("-----------\nNames and bdays:")
	for k, v := range bdays {
		fmt.Println(k, v)
	}
}
