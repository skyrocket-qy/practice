package atleastonce

import (
	"fmt"
	"time"
)

func makeRequestAtLeastOnce(url string, maxAttempts int) error {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := makeHTTPRequest(url)
		if err == nil {
			fmt.Println("HTTP request successful!")
			return nil
		}
		fmt.Printf("Attempt %d failed: %v\n", attempt, err)

		// Exponential backoff before the next attempt
		backoffDuration := time.Duration(1<<uint(attempt)) * time.Second
		time.Sleep(backoffDuration)
	}
	return fmt.Errorf("Exceeded maximum retry attempts")
}
