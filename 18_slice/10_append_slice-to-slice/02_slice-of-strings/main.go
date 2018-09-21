package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Webnesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}

// [Monday Tuesday Webnesday Thursday Friday]
