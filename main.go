package main

import (
	"fmt"

	"github.com/greyco1or/learngo2/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"
	definition, err := dictionary.Search("Second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}