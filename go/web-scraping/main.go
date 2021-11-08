package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const URL string = "https://cdkeys.com/daily-deals"

func handle_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", URL, nil)
	handle_error(err)

	req.Header.Set("User-Agent", "MyScraperBot v1.0")

	resp, err := client.Do(req)
	handle_error(err)
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	handle_error(err)
}
