package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/parnurzeal/gorequest"
	"github.com/toqueteos/webbrowser"
)

const htmlURL string = "http://kaifa.at/services/"
const jsonURL string = "http://kaifa.at/api/services/"

var (
	keyword = kingpin.Arg("keyword", "Keyword of service").Required().String()
	format  = kingpin.Flag("format", "Format: html,json").Default("html").String()
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()

	switch *format {
	case "html":
		openBrowser()
	case "json":
		fetchJSON()
	default:
		openBrowser()
	}
}

func openBrowser() {
	url := fmt.Sprint(htmlURL, *keyword)
	webbrowser.Open(url)
}

func fetchJSON() {
	url := fmt.Sprint(jsonURL, *keyword)

	request := gorequest.New()
	resp, body, errs := request.Get(url).
		Set("Accept", "application/json").
		End()

	if len(errs) > 0 {
		fmt.Fprintf(os.Stderr, "Unexpected errors: %s", errs)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Expected StatusCode=200, actual StatusCode=%v", resp.StatusCode)
		os.Exit(2)
	}

	fmt.Printf(body)
}
