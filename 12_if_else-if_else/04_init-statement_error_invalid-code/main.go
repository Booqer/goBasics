package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	/* food was initialized as part of the if statement,
	b is true so it prints food, which is chocolate
	commented out line below to remove error
	since it can't be accessed due to being block level scope */
	// fmt.Println(food)
}
