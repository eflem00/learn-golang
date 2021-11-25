package shared

import "log"

var something int

func init() {
	log.Println("one init called")
	something = 0
}

func Increment() {
	something = something + 1
	log.Printf("something: [%v]", something)
}