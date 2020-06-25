package main

import (
    "github.com/labstack/echo"
    "fmt"
    "encoding/json"
 )
func main() {
        e := echo.New()
e.GET("/get", func(c echo.Context) error {
        // data := c.QueryParam("message") #returning single value from query string
        queryString := c.QueryParams() // returning query object
       
        fmt.Println("Status: Get method Success..")
        fmt.Println("Query String is : ",queryString)
        return json.NewEncoder(c.Response()).Encode("Hello World.., This is Get Response")
        })

    e.POST("/post", func(c echo.Context) error {
        body_map := make(map[string]interface{})
        err := json.NewDecoder(c.Request().Body).Decode(&body_map)
    if err != nil {
        return err
    }
        fmt.Println("Status: Post Method Success..")
        fmt.Println("Body Params : ",body_map)
       return json.NewEncoder(c.Response()).Encode("Hello, All. This is Post method Response.")
    })
 e.Logger.Fatal(e.Start(":8080"))
}
