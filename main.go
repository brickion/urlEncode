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
      encoded, _ := UrlEncoded(scanner.Text())
      fmt.Println(encoded)
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
}


func UrlEncoded(str string) (string, error) {
  u, err := url.Parse(str)
  if err != nil {
      return "", err
  }
  return u.String(), nil
}
