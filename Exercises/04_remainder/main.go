package main

import (
	"fmt"
)

func main() {
	var bigNumba int
	var lilNumba int
	fmt.Println("Sup bro, tell me a big number and I'll do a trick...")
	fmt.Scan(&bigNumba)
	for bigNumba < 15 {
		fmt.Println("Do you even math, brah? You gotta know a number bigger than that.")
		fmt.Scan(&bigNumba)
	}
	fmt.Println("haha okay... now how big is your dong? No homo")
	fmt.Scan(&lilNumba)
	if lilNumba < 1 {
		fmt.Println("I was gonna pick on you, but that's pretty sad. You're sad. Or female, which scares me. Bye.")
	} else {
		thisMuch := bigNumba / lilNumba
		remainder := bigNumba % lilNumba
		fmt.Println("My dong is", thisMuch, "times as big as your's, with", remainder, "leftover")
		fmt.Println("Hah! scrub")
	}
}
