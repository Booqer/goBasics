package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"
	fmt.Println(myGreeting)

	var awShucks = map[string]string{}
	awShucks["Don't Fuckle"] = "With Shuckle"
	awShucks["Hi baby"] = "Hey Darlin"
	fmt.Println(awShucks)

	themApes := map[string]string{
		"Gorillas":    "Big, strong, and awesome",
		"Baboons":     "Butts reknowned across the globe",
		"Chimpanzees": "Smart and scary",
	}
	fmt.Println(themApes["Gorillas"])
	delete(themApes, "Baboons")
	themApes["Chimpanzees"] = "Actually quite friendly, let us play with your baby"
	fmt.Println(themApes)
}

// map[Tim:Good morning Jenny:Bonjour]
// map[Don't Fuckle:With Shuckle Hi baby:Hey Darlin]
// Big, strong, and awesome
// map[Gorillas:Big, strong, and awesome Chimpanzees:Actually quite friendly, let us play with your baby]
