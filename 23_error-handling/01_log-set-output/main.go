package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	// create a new file with a file name of log.txt and a variable name of nf
	nf, err := os.Create("log.txt")
	// if there is an error print it to the terminal
	if err != nil {
		fmt.Println(err)
	}

	// SetOutput takes a Writer interface, so passing it a pointer to a file satisifies this requirement
	// https://godoc.org/os#File.Write
	// https://godoc.org/io#Writer
	// https://godoc.org/log#SetOutput
	log.SetOutput(nf)
}

func main() {
	// no-file.txt doesn't exist, so it will throw an error
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		log.Println("err happened", err)
		// log.Fatalln(err)
		// panic(err)
	}
}

// Package log implements a simple logging package ... writes to standard error and prints the date
// and time of each logged message ... the Fatal functions call os.Exit(1) after writing the log message ..
// the Panic functions call panic after writing the log message

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println
