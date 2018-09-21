package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 5, 7, 9}
	myOtherSlice := []int{11, 13, 17, 19}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}

// [1 2 3 5 7 9 11 13 17 19]
