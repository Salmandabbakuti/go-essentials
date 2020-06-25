package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/thedevsaddam/gojsonq"
    "fmt"
)

func GetData(w http.ResponseWriter, r *http.Request) {

   jq := gojsonq.New().File("./response.json")
    res := jq.Find("meta.view.columns")

    if jq.Error() != nil {
        log.Fatal(jq.Errors())
    }
    fmt.Println("Server Responded To Get Method. Result is : ",res)
    json.NewEncoder(w).Encode(res)
}
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/get", GetData).Methods("GET")
    fmt.Println("Server Started Listening on Port 8000...")
    log.Fatal(http.ListenAndServe(":8000", router))
}
