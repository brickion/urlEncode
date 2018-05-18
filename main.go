package main

import (
  "bufio"
  "log"
  "os"
  "fmt"
  "net/url"
)

func main() {
  file, err := os.Open(os.Args[1])
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      encoded := url.QueryEscape(scanner.Text())
      fmt.Println(encoded)
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
}
