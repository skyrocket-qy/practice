package atmostonce

import (
	"fmt"
	"net/http"
)

func makeRequestAtMostOnce(url string, requestId string) error {
	// Include the unique identifier (requestId) in the request headers
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Request-ID", requestId)

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Process the response
	fmt.Println("HTTP request successful!")
	return nil
}
