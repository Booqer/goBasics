package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	// this is how you delete from a slice
	// append everything except that which you want deleted
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}

// [Monday Tuesday Wednesday Thursday Friday]
// [Monday Tuesday Thursday Friday]
