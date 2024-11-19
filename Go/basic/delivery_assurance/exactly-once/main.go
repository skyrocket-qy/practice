package main

import (
	"fmt"
)

func makeRequestExactlyOnce(url string) error {
	// Make an idempotent HTTP request (assuming the operation is idempotent)
	err := makeHTTPRequest(url)
	if err != nil {
		return err
	}
	fmt.Println("HTTP request successful (exactly once)!")
	return nil
}
