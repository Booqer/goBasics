package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
	// not exported values won't be Marshalled by json
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	// notice no 007
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

// notice no 007 in output:
//
// [123 34 70 105 114 115 116 34 58 34 74 97 109 101 115 34 44 34 76 97 115 116 34 58 34 66 111 110 100 34 44 34 65 103 101 34 58 50 48 125]
// []uint8
// {"First":"James","Last":"Bond","Age":20}
