package main

import "fmt"

func main() {
	theFinalCountdown := 1
	var masterList = []int{2}
	for i := 3; theFinalCountdown < 10001; i++ {
		for index, v := range masterList {
			if i%v == 0 {
				break
			} else if index == len(masterList)-1 {
				theFinalCountdown++
				masterList = append(masterList, i)
				fmt.Println(i, " looks to be a prime. Number: ", theFinalCountdown)
			}
		}
	}
	fmt.Println(masterList)
}

//  https://projecteuler.net/problem=7
//  find the 10001st prime number
