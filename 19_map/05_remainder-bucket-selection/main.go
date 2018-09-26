package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

// HashBucket takes the word and the number of buckets we want to create
// then returns the appropriate bucket for the word to go in
func HashBucket(word string, buckets int) int {
	// letter := rune(word[0])
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
