package main
import (
  "os"
  "io"
  "log"
  "fmt"
  "net/http"
  "encoding/json"
  "strconv"
  "posso-help/internal/chat"
  "posso-help/internal/db"
  "github.com/gorilla/mux"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
  log.Printf("HandleDownload()\n")
	vars := mux.Vars(r)
	datatype := vars["datatype"]
	phoneNumber := vars["phonenumber"]

  data, err := db.ReadOrdered(datatype, phoneNumber)
  if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		fmt.Fprintf(w, "%v", err)
		return 
  }

  csv, err := ConvertBsonToCsv(data) 
  if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		fmt.Fprintf(w, "%v", err)
		return 
  }

  length := strconv.Itoa(len(csv))
  disposition := fmt.Sprintf("attachment; filename=\"%s.csv\"", datatype)
  w.Header().Add("Content-Type", "text/html")
  w.Header().Add("Content-Length", length)
  w.Header().Add("Content-Disposition", disposition)
  fmt.Fprint(w, string(csv))

  return 
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
}

func HandleData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	datatype := vars["datatype"]
	phoneNumber := vars["phonenumber"]

  data, err := db.ReadUnordered(datatype, phoneNumber)
  if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		fmt.Fprintf(w, "%v", err)
		return 
  }

  json, err := json.Marshal(data)
  if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		fmt.Fprintf(w, "%v", err)
		return 
  }
  fmt.Fprint(w, string(json))
  return 
}

func HandleChatMessage(w http.ResponseWriter, r *http.Request) {
  log.Printf("HandleChatMessage")
  defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
    log.Printf("could not read body: %v\n", err)
		return
	}
	log.Printf("ChatMessage: %s\n", string(bodyBytes))

  chatMessage := &chat.ChatMessage{}
	err = json.Unmarshal(bodyBytes, chatMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("unmarshal error: %v\n", err)
		return
	}

  err = chat.ProcessEntries(chatMessage.Entries)
  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("error processing entries: %v\n", err)
		return
  }

  log.Printf("success handing chat message")
}

func HandleHubChallenge(w http.ResponseWriter, r *http.Request) {
  log.Printf("HandleHubChallenge")
  osToken := os.Getenv("HUB_TOKEN")
  mode := r.URL.Query().Get("hub.mode")
  token := r.URL.Query().Get("hub.verify_token")
  challenge := r.URL.Query().Get("hub.challenge")

  if len(osToken) < 1 {
    http.Error(w, "environment_error", http.StatusBadRequest)
		return
  }

  if mode != "subscribe" {
    http.Error(w, "invalid_mode", http.StatusBadRequest)
		return
	}

  if token != osToken {
    http.Error(w, "invalid_token", http.StatusBadRequest)
		return
	}

  fmt.Fprint(w, challenge)
}
