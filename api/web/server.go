package main

import (
    "net/http"
    "github.com/labstack/echo"
    "fmt"
    "encoding/json"
    "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
 )

  type Product struct {
    gorm.Model
    Code string
    Price string
    }

func main() {
        e := echo.New()
  e.GET("/getProduct", func(c echo.Context) error {
        Id := c.QueryParam("productId")
        db, err := gorm.Open("sqlite3", "apidata.db")
        queryString := c.QueryParams()
        if err != nil {
        panic("failed to connect database")
          }
       defer db.Close()
        var product Product
        db.First(&product, 1) // find product with id 1
        db.First(&product, "code = ?", Id)
        fmt.Println("Status: GetProduct Success..")
        fmt.Println("Query String is : ",queryString)
        fmt.Println("Response :",&product)
        return json.NewEncoder(c.Response()).Encode(product)
        })
   e.POST("/createProduct", func(c echo.Context) error {
        body_map := make(map[string]interface{})
        err := json.NewDecoder(c.Request().Body).Decode(&body_map)
    if err != nil {
        return err
     }else {
        price, ok := body_map["price"].(string)
        productId := body_map["productId"].(string)
        db, err := gorm.Open("sqlite3", "apidata.db")
       if err != nil {
        panic("failed to connect database")
          }
       defer db.Close()
       db.AutoMigrate(&Product{})

       db.Create(&Product{Code:productId, Price:price})

        fmt.Println("Status: Create Product Success..")
        fmt.Println("Body Params : ",body_map)
       fmt.Println(ok)
       return c.String(http.StatusOK, "Product Added.")
       }
    })
 e.Logger.Fatal(e.Start(":8080"))
}
