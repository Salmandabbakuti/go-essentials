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
            //values := c.QueryParams()
        //fmt.Println(value)
        //fmt.Println(values)
        db, err := gorm.Open("sqlite3", "apidata.db")
        if err != nil {
        panic("failed to connect database")
          }
       defer db.Close()
     db.First(&product, 1) // find product with id 1
    db.First(&product, "code = ?", Id)
    var data Product
    err := json.NewDecoder(product).Decode(&data)
  if err != nil {
      fmt.Println(err)
}
    //data:= json.Unmarshal(product, Product)
    fmt.Println(&product)
    fmt.Println("GetProduct Success..")
    return c.String(http.StatusOK, data)
    })
    e.POST("/createProduct", func(c echo.Context) error {

     body_map := make(map[string]interface{})
    err := json.NewDecoder(c.Request().Body).Decode(&body_map)
    if err != nil {
     return err
} else {
    //json_map has the JSON Payload decoded into a map
    price, ok := body_map["price"].(string)
   // p, err := strconv.ParseUint(price,16,32)
    // v :=uint(p)
            code := body_map["code"].(string)
    db, err := gorm.Open("sqlite3", "apidata.db")
        if err != nil {
        panic("failed to connect database")
          }
       defer db.Close()
       db.AutoMigrate(&Product{})

    db.Create(&Product{Code: code, Price:price})
     fmt.Println("Data Posted..")
     fmt.Println(ok)
     fmt.Println(body_map)
      return c.String(http.StatusOK, "Product Added.")
         }
        })
        e.Logger.Fatal(e.Start(":8080"))
}
