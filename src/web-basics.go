package main

import ("fmt"
        "net/http")

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1> hello....</h1>")
  fmt.Fprintf(w, "<p> hello....</p>")
  fmt.Fprintf(w, "<p> hello....</p>")
  fmt.Fprintf(w, "<p> hello %s</p>", "<strong>can even</strong>")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":8000", nil)
}
