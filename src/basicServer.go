package main

/**
To run this 
go run basicServer.go

curl http://localhost:8000
curl http://localhost:8000/about

**/
import ("fmt"
        "net/http")

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whoa, Go is neat so far")
}

func aboutHandler(w http.ResponseWriter, r * http.Request) {
  fmt.Fprintf(w, "About page")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/about", aboutHandler)
  http.ListenAndServe(":8000", nil)
}
