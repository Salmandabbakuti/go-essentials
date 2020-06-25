package main

import (
    "net/http"
    "github.com/labstack/echo"
    "fmt"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "github.com/thedevsaddam/gojsonq"
    "io/ioutil"
    //"github.com/go-resty/resty"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
 )

  type Data struct {
    gorm.Model
    DiseaseId int
    DiseaseName string
    Year int
    TotalCount int
    State string
     
    }

func main() {
        e := echo.New()
  e.GET("/getData", func(c echo.Context) error {
        Id := c.QueryParam("queryId")
        db, err := gorm.Open("sqlite3", "apidata.db")
        queryString := c.QueryParams()
        if err != nil {
        panic("failed to connect database")
          }
       defer db.Close()
        var data Data
        db.First(&data, 1) // find data with id 1
        db.First(&data, "Id = ?", Id)
        fmt.Println("Status: Getdata Success..")
        fmt.Println("Query String is : ",queryString)
        fmt.Println("Response :",&data)
        return json.NewEncoder(c.Response()).Encode(data)
        })
   e.POST("/createData", func(c echo.Context) error {
        //body_map := make(map[string]interface{})
        //err := json.NewDecoder(c.Request().Body).Decode(&body_map)
        //client := resty.New()
        //resp, err := client.R().EnableTrace().Get("https://data.cdc.gov/api/views/dwqk-w36f/rows.json?accessType=DOWNLOAD")
        response, err := http.Get("https://data.cdc.gov/api/views/dwqk-w36f/rows.json?accessType=DOWNLOAD")
        data, _ := ioutil.ReadAll(response.Body)
        //resp := gojsonq.New().FromString(string(data)).Find("data")
      if err != nil {
        return err
     }else {
         
         for i := 0; i < 100; i++ {
        diseaseName:= gojsonq.New().FromString(string(data)).Find("data[i].8")
        year:= gojsonq.New().FromString(string(data)).Find("data[i].9")
        totalCount:= gojsonq.New().FromString(string(data)).Find("data[i].10")
        state:= gojsonq.New().FromString(string(data)).Find("data[i].27")
        db, err := gorm.Open("sqlite3", "apidata.db")
       if err != nil {
        panic("failed to connect database")
          }
       defer db.Close()
       //db.AutoMigrate(&Product{})
       db.Create(&Data{Id:i, DiseaseName:diseaseName.(string), Year:year.(int), TotalCount:totalCount.(int), State:state.(string)})
	           }
       fmt.Println("Status: Create Product Success..")
       return c.String(http.StatusOK, "Products Added.")
        }
    })
 e.Logger.Fatal(e.Start(":8080"))
}
