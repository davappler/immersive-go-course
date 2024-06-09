package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	url := "http://localhost:8080"
	max_tries := 5

	fmt.Println("----------------------------")

	for i := 0; i < max_tries; i++ {
		resp, err := c.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error from request:", err)
			return
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfter := resp.Header.Get("Retry-After")
			retryAfterInteger, err := strconv.Atoi(retryAfter)
			if err != nil {
				// This is where we give up if we do not have the amount of time as an integer. We would not like to wait for it as it is uncertain how long is it gonna be.
				fmt.Fprintln(os.Stderr, "Error parsing string", err)
				return
			}

			if retryAfterInteger > 1 {
				fmt.Println("There is a little delay...")

			}

			if retryAfterInteger > 5 {
				fmt.Fprintln(os.Stderr, "Server is overloaded we can not get the weather details sorry", err)
				return
			}

			duration := time.Duration(retryAfterInteger) * time.Second
			time.Sleep(duration)
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error reading response body", err)
				return
			}
			fmt.Fprintln(os.Stdout, string(body))
			return
		}
	}

	fmt.Fprintln(os.Stderr, "Max tries reached! Exiting")
}
