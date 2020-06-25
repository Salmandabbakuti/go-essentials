package main

import (
    "github.com/go-resty/resty"
    "fmt"
      )

func main() {
        // GET request
  client := resty.New()

  resp, err := client.R().
  SetHeader("Content-Type", "application/json").
  SetBody(`{"username":"testuser", "password":"testpass"}`).Post("http://localhost:8080/post")
  fmt.Println(resp)
  fmt.Println(err)
}
