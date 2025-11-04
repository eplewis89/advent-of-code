package bootstrap

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://adventofcode.com/%d/day/%d/input"

func ReadInput(year int, day int) (string, error) {
	url := fmt.Sprintf(baseURL, year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", fmt.Errorf("error creating http request: %v", err)
	}

	// Use session token from environment variable
	req.Header.Add("Cookie", "session="+os.Getenv("AOC_SESSION"))

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error making http request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected http status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}
