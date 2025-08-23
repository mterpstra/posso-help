
package main

import (
  "io"
  "testing"
  "net/http/httptest"

  "github.com/gorilla/mux"
  "github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
  router:= mux.NewRouter()
  router.HandleFunc("/data/download/{datatype}/{phonenumber}", 
    HandleDownload)
  req := httptest.NewRequest("GET", "/data/download/births/16166100305", nil)
	rsp := httptest.NewRecorder()
	router.ServeHTTP(rsp, req)
  httpResponse := rsp.Result()
  bytes, err := io.ReadAll(httpResponse.Body)
  assert.Nil(t, err, "error was not nil")
  assert.Greater(t, len(string(bytes)), 10, "not enough bytes returned")
  assert.Equal(t, 
    httpResponse.Header.Get("Content-Disposition"),
    "attachment; filename=\"births.csv\"",
    "content dispossition header wrong")
}

func TestDataFetch(t *testing.T) {
  router:= mux.NewRouter()
  router.HandleFunc( "/data/{datatype}/{phonenumber}", HandleData)
  req := httptest.NewRequest( "GET", "/data/births/16166100305", nil)
	rsp := httptest.NewRecorder()
	router.ServeHTTP(rsp, req)
  httpResponse := rsp.Result()
  bytes, err := io.ReadAll(httpResponse.Body)
  assert.Nil(t, err, "error was not nil")
  assert.Greater(t, len(string(bytes)), 10, "not enough bytes returned")
  assert.Equal(t, 
    "text/plain; charset=utf-8",
    httpResponse.Header.Get("Content-Type"),
    "content-type header wrong")
}
