package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	jdata  Stations
	client = &http.Client{
		Timeout: 2 * time.Second,
	}
	handle_error = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	ndata Stations
)

type Station struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Stations struct {
	Values []Station `json:"stations"`
}

func init() {
	jf, err := os.Open("/home/z1gzag/projects/radio.json")
	handle_error(err)
	defer jf.Close()
	bytes, _ := ioutil.ReadAll(jf)

	err = json.Unmarshal(bytes, &jdata)
	handle_error(err)
}

func main() {
	for i, d := range jdata.Values {
		resp, err := req(d.Link)
		fmt.Println("Current index: ", i)

		if err != nil {
			fmt.Printf("Error encountered: %v\n", err)
			continue
		}

		if resp.StatusCode == 200 {
			ndata.Values = append(ndata.Values, d)
		}
	}

	nd, _ := json.Marshal(ndata)
	err := ioutil.WriteFile("ease.json", nd, 0644)
	handle_error(err)

	fmt.Println("Wrote some ease!")
}

func req(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	handle_error(err)
	req.Header.Set("User-Agent", "MyScraperBot v1.0")

	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
	}

	return resp, err
}
