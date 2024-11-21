package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const url = "http://localhost:8080/"

func main() {
	err := getWeather(0)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func getWeather(maxAttempt int) error {
	if maxAttempt >= 3 {
		return fmt.Errorf("max attempts allowed reached : %v", maxAttempt)
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http request or connection failed  or dropped: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		retryTime, err := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			return fmt.Errorf("error reading a retry time: %v", err)
		}

		if retryTime > 5 {
			return fmt.Errorf("server request too long delay time : %d seconds", retryTime)
		}

		fmt.Fprintf(os.Stderr, "system is a bit slow we're trying again in: %v seconds\n", retryTime)
		time.Sleep(time.Duration(retryTime) * time.Second)
		err = getWeather(maxAttempt + 1)
		if err != nil {
			return fmt.Errorf("an error occured trying to retry the request: %v", err)
		}
	}

	if resp.StatusCode == 200 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response: %v", err)
		}
		fmt.Fprintf(os.Stdout, "Results for the weather: %v", string(msg))
		return nil
	}

	return fmt.Errorf("request returned an invalid status code: %v", resp.StatusCode)
}
