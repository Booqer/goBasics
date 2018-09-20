package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Derek"])
}

func changeMe(z map[string]int) {
	z["Todd"] = 44
}
