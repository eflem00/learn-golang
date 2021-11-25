package main

import "fmt"

func main() {
	// init 1
	colors := map[string]string{
		"red": "#ff0000",
	}
	fmt.Println(colors["test"])

	//init 2
	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println(colors2["white"])

	// delete
	delete(colors2, "white")
	fmt.Println(colors2["white"])

	//iterate
	colors3 := map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}
	for key, val := range colors3 {
		fmt.Println(fmt.Sprintf("k: %v v: %v", key, val))
	}
}
