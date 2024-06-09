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
			os.Exit(1) // Exit code 1 for network or request error
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfter := resp.Header.Get("Retry-After")

			if retryAfter == "a while" {
				fmt.Fprintln(os.Stderr, "Server is overloaded we can not get the weather details sorry", err)
				os.Exit(3) // Exit code 3 for server overload
			}

			// here I can check I am getting a timestamp and find how far in time is that timestamp, if it greater than a the threshold we return exit.code(3) (I need to decide the timestamp)
			retryAfterInteger, err := strconv.Atoi(retryAfter)
			if err != nil {
				// This is where we give up if we do not have the amount of time as an integer. We would not like to wait for it as it is uncertain how long is it gonna be.
				fmt.Fprintln(os.Stderr, "Error parsing string", err)
				os.Exit(2) // Exit code 2 for parsing error
			}

			if retryAfterInteger > 5 {
				fmt.Fprintln(os.Stderr, "Server is overloaded we can not get the weather details sorry", err)
				os.Exit(3) // Exit code 3 for server overload
			} else if retryAfterInteger > 1 {
				fmt.Println("There is a little delay...")

			}

			duration := time.Duration(retryAfterInteger) * time.Second
			time.Sleep(duration)
		} else {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error reading response body", err)
				os.Exit(4) // Exit code 4 for reading response body error
			}
			fmt.Fprintln(os.Stdout, string(body))
			return
		}
	}

	fmt.Fprintln(os.Stderr, "Max tries reached! Exiting")
	os.Exit(5) // Exit code 5 for max tries reached
}
