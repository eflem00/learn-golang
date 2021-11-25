package main

import (
	"example/how-to/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	str := morestrings.ReverseRunes("!oG ,olleH")

	fmt.Println(str)

	fmt.Println(morestrings.Other("something"))

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}