package main
import (
  "os"
  "io"
  "log"
  "fmt"
  "time"
  "net/http"
  "encoding/json"
  "context"
  "strconv"
  "posso-help/internal/chat"
  "posso-help/internal/db"
  "posso-help/internal/user"
  "github.com/gorilla/mux"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
  log.Printf("HandleDownload()\n")
	vars := mux.Vars(r)
	datatype := vars["datatype"]

  ctx := r.Context()
  userID := ctx.Value("user_id")
  if userID == nil {
    log.Printf("could not get userid from context")
    http.Error(w, "Authorization header required", http.StatusUnauthorized)
    return
  }

  user, err := user.Read(userID.(string))
  if err != nil {
    log.Printf("could not read userID from context")
    http.Error(w, "User Not Found", http.StatusNotFound)
    return
  }

  data, err := db.ReadOrdered(datatype, user.PhoneNumbers)
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

func HandleDataGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	datatype := vars["datatype"]

  ctx := r.Context()
  userID := ctx.Value("user_id")
  if userID == nil {
    log.Printf("could not get userid from context")
    http.Error(w, "Authorization header required", http.StatusUnauthorized)
    return
  }

  user, err := user.Read(userID.(string))
  if err != nil {
    log.Printf("could not read userID from context")
    http.Error(w, "User Not Found", http.StatusNotFound)
    return
  }

  data, err := db.ReadUnordered(datatype, user.PhoneNumbers)
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

func HandleDataPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	datatype := vars["datatype"]
  // @todo: check valid values for datatype 
  log.Printf("HandleDataPost: %s", datatype)
  ctx := r.Context()
  userID := ctx.Value("user_id")
  if userID == nil {
    log.Printf("could not get userid from context")
    http.Error(w, "Authorization header required", http.StatusUnauthorized)
    return
  }

  u, err := user.Read(userID.(string))
  if err != nil {
    log.Printf("could not read userID from context")
    http.Error(w, "User Not Found", http.StatusNotFound)
    return
  }

  collection := db.GetCollection(datatype);

  defer r.Body.Close()
  bodyBytes, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "Error reading request body", http.StatusInternalServerError)
    log.Printf("Error reading request body: %v", err)
    return
  }

  log.Printf("user: %s  collection: %v  body: %s",
             u.Username, collection, string(bodyBytes))

  data := make(map[string]interface{})

  err = json.Unmarshal(bodyBytes, &data)
  if err != nil {
    http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
    log.Printf("Error unmarshalling JSON: %v", err)
    return 
  }

  data["entry_id"]   = fmt.Sprintf("dashboard-%d", time.Now().Unix())
  data["message_id"] = "manual-entry"
  data["phone"]      = u.PhoneNumbers[0]
  data["name"]       = u.Name
  if data["name"] == "" {
    data["name"] = u.Username
  }

  _, err = collection.InsertOne(context.TODO(), data)
  if err != nil {
    http.Error(w, "Error Inserting Temperature", http.StatusBadRequest)
    log.Printf("Error Inserting Temperature: %v", err)
    return 
  }
  
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
