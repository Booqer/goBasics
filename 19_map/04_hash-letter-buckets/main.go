package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	// defer will close the response body right before the function exits
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	// fmt.Println("**********************")
	// for i := 28; i <= 126; i++ {
	// fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	// }
}

// HashBucket converts the first letter of a passed word into an integer and returns the int
func HashBucket(word string) int {
	return int(word[0])
}
