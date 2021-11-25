package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func getHealth(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New("Bad status code")
	}

	return nil
}

func processUrl(url string, c chan string) {
	err := getHealth(url)

	if err != nil {
		fmt.Println(fmt.Sprintf("ERR %v", url))
	} else {
		fmt.Println(fmt.Sprintf("OK %v", url))
	}

	c <- url
}

// using channels
func main() {
	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
	}

	channel := make(chan string)

	for _, site := range sites {
		go processUrl(site, channel)
	}

	// kind of special syntax, iterate over the channel indefinitely listening for messages
	for url := range channel {
		go func(url string) {
			// wait a bit
			time.Sleep(time.Second)
			processUrl(url, channel)
		}(url)
	}
}

// using waitgroup
// func main2() {
// 	sites := []string{
// 		"https://google.com",
// 		"https://facebook.com",
// 		"https://stackoverflow.com",
// 	}

// 	var wg sync.WaitGroup

// 	for i, site := range sites {
// 		wg.Add(i)

// 		go func(site string) {
// 			defer wg.Done()
// 			processUrl(site)
// 		}(site)
// 	}

// 	wg.Wait()

// 	fmt.Println("Done my work...")
// }
