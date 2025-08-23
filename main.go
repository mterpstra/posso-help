package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"os"
	"strings"

  "github.com/gorilla/mux"
)

func logRequest(r *http.Request) {
	uri := r.RequestURI
	method := r.Method
	fmt.Println("Got request!", method, uri)
}

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}


func main() {
	r := mux.NewRouter()

	// Apply the logging middleware to all routes
	r.Use(LoggingMiddleware)

  // Serve static files from the "static" directory
  staticFileDirectory := http.Dir("./static/")
  fileServer := http.FileServer(staticFileDirectory)
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! you've requested %s\n", r.URL.Path)
	})

	r.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		keys, ok := r.URL.Query()["key"]
		if ok && len(keys) > 0 {
			fmt.Fprint(w, os.Getenv(keys[0]))
			return
		}
		envs := []string{}
		envs = append(envs, os.Environ()...)
		fmt.Fprint(w, strings.Join(envs, "\n"))
	})

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `{"status":"success"}`)
	})

  log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":80", r))
}
