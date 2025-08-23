package main
import (
  "log"
  "fmt"
  "net/http"
  "encoding/json"
  "strconv"

  "github.com/gorilla/mux"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
  log.Printf("HandleDownload()\n")
	vars := mux.Vars(r)
	datatype := vars["datatype"]
	phoneNumber := vars["phonenumber"]

  data, err := dbReadOrdered(datatype, phoneNumber)
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

  data, err := dbReadUnordered(datatype, phoneNumber)
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
