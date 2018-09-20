package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	// fmt.Println(sf)
	// sf is [43, 56, 87, 12, 45, 57]
	// fmt.Printf("%T \n", sf)
	// shows []float64
	// which is a slice of float64
	var total float64
	// range sf will loop over everything in the slice
	// it will return the index and the value, we don't need
	// the index, only the value, so we throw away the index
	// with the blank identifier
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
