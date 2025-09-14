package main

import (
  "fmt"
  "time"
  "log"
  "net/http"
  "os"
  "strings"
  "sync"

  "github.com/gorilla/mux"
)

var mu sync.Mutex
var termColors []string
var colorIndex int
func init() {
  colorIndex = 0
  termColors = []string {
    "\033[30m",
    "\033[31m",
    "\033[32m",
    "\033[33m",
    "\033[34m",
    "\033[35m",
    "\033[36m",
    "\033[37m",
  }
}

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // @todo: Remove this mutex for production builds.  This forces
    //        a single request to finish before the next starts.
    //        Nice for debugging since React in development mode 
    ///       fires two useEffect hooks to force strong components.
    mu.Lock()
    defer mu.Unlock()

    if colorIndex >= len(termColors) {
      colorIndex = 0;
    }
    print(termColors[colorIndex])
    start := time.Now()
    log.Printf("Started %s %s", r.Method, r.URL.Path)
    next.ServeHTTP(w, r) // Call the next handler in the chain
    log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
    print("\033[0m")
    colorIndex++;

  })
}

func main() {
  r := mux.NewRouter()

  // Apply the logging middleware to all routes
  r.Use(LoggingMiddleware)


  // Non-Auth register/login routes
  r.HandleFunc("/api/auth/register", HandleAuthRegister).Methods("POST")
  r.HandleFunc("/api/auth/verify-email", HandleEmailVerification)
  r.HandleFunc("/api/auth/login", HandleLogin)


  // Handle the webhook setup request
  r.HandleFunc("/chat/message", HandleHubChallenge).
  Methods("GET").
  Queries("hub.mode", "{mode}").
  Queries("hub.verify_token", "{token}").
  Queries("hub.challenge", "{challenge}")

  // Text Message Handler
  r.HandleFunc("/chat/message", HandleChatMessage).Methods("POST")

  // Data routes
  dataRouter := r.PathPrefix("/api/data").Subrouter()
  dataRouter.Use(AuthMiddleware) 
  dataRouter.HandleFunc("/{datatype}", HandleDataGet).Methods("GET")
  dataRouter.HandleFunc("/births", HandleBirthPost).Methods("POST")
  dataRouter.HandleFunc("/deaths", HandleDeathPost).Methods("POST")
  dataRouter.HandleFunc("/temperature", HandleTemperaturePost).Methods("POST")
  dataRouter.HandleFunc("/rain", HandleRainPost).Methods("POST")

  // Download routes
  downloadRouter := r.PathPrefix("/api/download").Subrouter()
  downloadRouter.Use(AuthMiddleware) 
  downloadRouter.HandleFunc("/{datatype}/{phonenumber}", HandleDownload).Methods("GET")

  userRouter := r.PathPrefix("/api/user").Subrouter()
  userRouter.Use(AuthMiddleware) 
  userRouter.HandleFunc("/phonenumber", HandleLinkPhoneNumber).Methods("POST")
  userRouter.HandleFunc("", HandleGetUser).Methods("GET")

  // Serve static files from the "static" directory
  staticFileDirectory := http.Dir("./static/")
  fileServer := http.FileServer(staticFileDirectory)
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

  // Old Stuff
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello! you've requested %s\n", r.URL.Path)
  })

  r.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
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
  log.Fatal(http.ListenAndServe(":8080", r))
}
