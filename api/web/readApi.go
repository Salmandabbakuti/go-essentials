
package main

import (
        "github.com/go-resty/resty"
        "fmt"
)

func main() {
        // GET request
client := resty.New()
resp, err := client.R().EnableTrace().Get("http://localhost:8080/get")

// explore response object
fmt.Printf("\nError: %v", err)
fmt.Printf("\nResponse Status Code: %v", resp)
}
