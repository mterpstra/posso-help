package main

import (
  "testing"
  "os"
	"math/rand"
	"time"
  "strings"
  "bytes"
  "net/http"
  "io"
  "fmt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return "test-" + string(b)
}

func replaceControlChars(s string) string {
	var result bytes.Buffer
	for _, r := range s {
    if r == 0x0a {
			result.WriteString("\\n")
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

/*
func TestMainHubRequest(t *testing.T) {
  token := "some-set-of-bytes"
	os.Setenv("HUB_CHALLENGE_TOKEN", token)
  input := map[string]interface{}{}
  input["hub.challenge"] = "token"
  input["hub.mode"] = "mode"
  input["hub.verify_token"] = token
  response := Main(input)
  if response["statusCode"].(int) != 200 {
    t.Errorf("Main return error [%d] [%v]", response["statusCode"].(int), response["body"])
  }
}
*/

func TestHandleEntryRequest(t *testing.T) {
	template, err := os.ReadFile("./test_data/test.json")
	if err != nil {
		t.Errorf("Error reading test template: %v", err)
  }

  dirPath := "./test_data"
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		t.Errorf("Error reading directory: %v", err)
	}

	for _, entry := range entries {

    if !strings.HasSuffix(entry.Name(), "txt") {
      continue
    }

    raw, err := os.ReadFile("./test_data/" + entry.Name())
    if err != nil {
      t.Errorf("Error reading test data: %v", err)
    }

    fileData := replaceControlChars(string(raw))
    testData := string(template)
    testData = strings.Replace(testData, "{message_id}", generateRandomString(50), -1)
    testData = strings.Replace(testData, "{event_id}",   generateRandomString(20), -1)
    testData = strings.Replace(testData, "{message_body}", fileData, -1) 
    
    println("-------------------------------------")
    println("integration test", entry.Name())
    println(testData)

    url := "http://localhost:8080/chat/message"
    resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(testData)))
    if err != nil {
      t.Errorf("error sending data to server")
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
      t.Errorf("error reading data")
    }
    fmt.Printf("server response: %s\n", string(body))

    println("-------------------------------------")
	}
}
