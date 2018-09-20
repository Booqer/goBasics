package main

import "fmt"

func main() {
	age := 34
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 12
	fmt.Println(z)
	fmt.Println(*z)
}
