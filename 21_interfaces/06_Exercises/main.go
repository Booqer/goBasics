/*
	use https://godoc.org/sort to sort out the following:

	(1)
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	(2)
	s := []string{"Zeno", "John", "Al", "Jenny"}

	(3)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	sort each in more than one way, and then also in reverse
*/

package main

import (
	"fmt"
	"sort"
)

type people []string

var studyGroup = people{"Zeno", "John", "Al", "Jenny"}

var spookHouse = []string{"Fang", "Kaitlin", "Krystle", "Derek"}

var justNums = []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

func main() {
	fmt.Println("studyGroup: ", studyGroup)
	sort.Slice(studyGroup, func(i, j int) bool { return studyGroup[i] < studyGroup[j] })
	fmt.Println("studyGroup sorted: ", studyGroup)
	// sort.Reverse(sort.Slice(studyGroup, func(i, j int) bool { return studyGroup[i] < studyGroup[j] }))

	fmt.Println("spookHouse: ", spookHouse)
	sort.Strings(spookHouse)
	fmt.Println("spookHouse sorted: ", spookHouse)
	// sort.Sort(sort.Reverse(sort.Strings(spookHouse)))

	fmt.Println("justNums: ", justNums)
	sort.Ints(justNums)
	fmt.Println("justNums sorted: ", justNums)
	sort.Sort(sort.Reverse(sort.IntSlice(justNums)))
	fmt.Println("justNums reversed: ", justNums)
}
