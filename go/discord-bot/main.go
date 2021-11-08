package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Prefix string `json:"prefix"`
	Devs   []int  `json:"devs"`
}

var (
	TOKEN   string = ""
	VERBOSE bool   = false
	CONFIG  Config
)

func handle_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	flag.StringVar(&TOKEN, "t", "", "Used to interact with the Discord API.")
	flag.Parse()

	jf, err := os.Open("config.json")
	handle_error(err)
	defer jf.Close()
	bytes, _ := ioutil.ReadAll(jf)

	err = json.Unmarshal(bytes, &CONFIG)
	handle_error(err)
}

func main() {
	if VERBOSE {
		fmt.Println("Using verbose logging.")
	}

}
