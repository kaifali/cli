package main

import (
  "fmt"

  "github.com/toqueteos/webbrowser"
  "github.com/alecthomas/kingpin"
)

var (
  keyword = kingpin.Arg("keyword", "Keyword of service").Required().String()
)

func main() {
  kingpin.Version("0.0.2")
  kingpin.Parse()

  url := fmt.Sprint("http://kaifa.li/services/", *keyword)
  webbrowser.Open(url)
}
