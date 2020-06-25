
package main
import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "log"
    "github.com/thedevsaddam/gojsonq"
    "fmt"
    "io/ioutil"
    "net/http"
)
func GetData(w http.ResponseWriter, r *http.Request) {
   response, err := http.Get("https://data.cdc.gov/api/views/dwqk-w36f/rows.json?accessType=DOWNLOAD")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        resp := gojsonq.New().FromString(string(data)).Find("data")
       fmt.Println("Server Responded To Get Method..")
       json.NewEncoder(w).Encode(resp)
    }
}
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/get", GetData).Methods("GET")
    fmt.Println("Server Started Listening on Port 8000...")
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
