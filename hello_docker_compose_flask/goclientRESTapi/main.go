package main

import (
//   "bytes"
//    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    fmt.Println("Starting the application...")
    response, err := http.Get("http://observatory-service/")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println("Sahil DATA:",string(data))
    }
    
    fmt.Println("SahilRESTGoClient:Terminating the application...")
}
