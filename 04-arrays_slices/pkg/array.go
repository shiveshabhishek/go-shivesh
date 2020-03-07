package array

import (
	"errors"
	"math/rand"
)

const size = 5 // size of the array

// Create array with random Entries
func CreateArrayWithEntries() [size]int {
	arr := [size]int{}
	for i := 0; i < size; i++ {
		arr[i] = (rand.Intn(100)) // fill array with random values
	}
	return arr
}

// find minimum value amongst all entries in the array
func MaxVal(arr [size]int) (int,error) {
	var MAX = -1
	if size==0{
		return 0, errors.New("Empty Array")
	}
	for i := 0; i < size; i++ {
		if arr[i] > MAX {
			MAX = arr[i]
		}
	}
	return MAX,nil
}

// find minimum value amongst all entries in the array
func MinVal(arr [size]int) (int,error) {
	MIN := 200
	if size==0{
		return 0, errors.New("Empty Array.")
	}
	for i := 0; i < size; i++ {
		if arr[i] < MIN {
			MIN = arr[i]
		}
	}
	return MIN,nil
}
