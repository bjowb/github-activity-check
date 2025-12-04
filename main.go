package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type GitHubEvent struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	Repo      struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func fetchEvents(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []GitHubEvent
	err = json.NewDecoder(resp.Body).Decode(&events)
	return events, nil
}

func main() {
	username := flag.String("user", "", "Github Username ")

	flag.Parse()

	if *username == "" {
		fmt.Println("Error: --user <username> is required")
		flag.Usage()
		os.Exit(1)
	}

	events, err := fetchEvents(*username)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("recent github activities for %s\n\n", *username)

	for i, e := range events {
		t, _ := time.Parse(time.RFC3339, e.CreatedAt)
		fmt.Printf("%2d. %-18s %-30s %s\n", i+1, e.Type, e.Repo.Name, t.Format("2006-01-02 15:04"))
		if i == 15 {
			break
		}
	}
}
