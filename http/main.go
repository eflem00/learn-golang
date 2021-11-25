package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Printf("got here: %v", string(bs))
	return len(bs), nil
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error calling google.com")
		os.Exit(1)
	}

	// body to string
	// defer resp.Body.Close()

	// body := make([]byte, 99999)
	// _, err = resp.Body.Read(body)

	// fmt.Println(string(body))

	// body to json
	// var result interface{}
	// err = json.NewDecoder(resp.Body).Decode(&result)

	// if err != nil {
	// 	fmt.Printf("error parsing body: %v", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%#v", result)

	// body to custom writer
	lw := logWriter{}
	num, err := io.Copy(lw, resp.Body)

	fmt.Printf("result: %v %v", num, err)
}
