package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s!", r.URL)
    fmt.Println("RESTfulServ. on:80, Controller:",r.URL.Path[1:])
}
func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting Restful services...")
    fmt.Println("Using port:80")
    err := http.ListenAndServe(":80", nil)
    log.Print(err)
    errorHandler(err)
}
func errorHandler(err error){
if err!=nil {
    fmt.Println(err)
    os.Exit(1)
}
}
