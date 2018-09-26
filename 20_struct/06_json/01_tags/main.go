package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	// the - tag means don't include this
	Last string `json:"-"`
	Age  int    `json:"wisdom score"`
}

func main() {
	p1 := person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

// notice no last name in output:
// {"First":"James","wisdom score":20}
