package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	if len(os.Args) < 2 {
		handleError(fmt.Errorf("usage: github-activity <username>"))
	}

	userName := os.Args[1]
	url := GenerateUrl(userName)

	response, err := FetchData(url)

	if err != nil {
		handleError(fmt.Errorf("error Fetching data : %s", err))
	}

	// آنمارشال کردن داده‌ها
	events, decodeErr := DecodeEvents(response)
	if decodeErr != nil {
		handleError(decodeErr)
	}

	DisplayEvents(events)
}
func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("An unknown error occurred.")
	}
	os.Exit(1)
}

func GenerateUrl(userName string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events", userName)
}

func FetchData(url string) (io.Reader, error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error: %s \n", err)
	}
	if res == nil {
		return nil, fmt.Errorf("received a nil response")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: Received status code %d\n", res.StatusCode)
	}

	// Convert the body back to an io.Reader to pass to DecodeEvents
	return res.Body, nil
}

func DecodeEvents(reader io.Reader) ([]Event, error) {
	var events []Event
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&events); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}
	return events, nil
}

func DisplayEvents(events []Event) {
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("Pushed to %s\n", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("Created a pull request in %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("Opened an issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Starred %s\n", event.Repo.Name)
		default:
			fmt.Printf("Other activity in %s\n", event.Repo.Name)
		}
	}
}
