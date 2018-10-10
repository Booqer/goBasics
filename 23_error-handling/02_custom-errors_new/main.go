package main

import (
	// we are importing the errors package
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Norgate math: square root of negative number")
	}
	// implementation
	return 42, nil
}
