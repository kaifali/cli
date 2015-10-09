package main

import "os"
import "fmt"
import "github.com/toqueteos/webbrowser"

func main() {
  keyword := os.Args[1]

  url := fmt.Sprint("http://kaifa.li/services/", keyword)
  webbrowser.Open(url)
}
