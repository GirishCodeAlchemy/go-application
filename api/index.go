package handler

import (
	"log"
	"net/http"
)

// func main() {
//     fs := http.FileServer(http.Dir("static"))
//     http.Handle("/", http.StripPrefix("/", fs))

//     log.Println("Server started on :8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }


func Handler(w http.ResponseWriter, r *http.Request) {
  // Serve static files if the URL path starts with "/static/"
  if r.URL.Path == "/" {
    // fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
//   } else if r.URL.Path == "/" || len(r.URL.Path) > 8 && r.URL.Path[:8] == "/static/" {
    fs := http.FileServer(http.Dir("static"))
    http.StripPrefix("/", fs).ServeHTTP(w, r)
  } else {
    http.NotFound(w, r)
  }
}


