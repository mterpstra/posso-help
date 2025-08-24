package textmsg

import (
	"bytes"
	"encoding/json"
  "os"
	"log"
	"fmt"
	"errors"
	"net/http"
)

var authToken, phoneNumberID, url string
func init() {
  authToken = os.Getenv("AUTH_TOKEN")
  phoneNumberID = os.Getenv("PHONE_NUMBER_ID")

  if authToken == "" {
    log.Printf("Missing AUTH_TOKEN needed for sending text replys")
  }

  if phoneNumberID == "" {
    log.Printf("Missing PHONE_NUMBER_ID needed for sending text replys")
  }

  url = fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", phoneNumberID)
}

type Text struct {
  Body string `json:"body"`
}

type MessageSender struct {
	MsgProd string `json:"messaging_product"`
	Type    string `json:"type"`
	To      string `json:"to"`
	Text    Text   `json:"text"`
}

func NewMessageSender(to, body string) *MessageSender {
  return &MessageSender {
    MsgProd: "whatsapp",
    Type: "text",
    To: to,
    Text: Text {
      Body: body,
    },
  }
}

func (ms *MessageSender) getPayload() ([]byte, error) {
	return json.Marshal(ms)
}

func (ms *MessageSender) validate() error {
  if (authToken == "") {
    return errors.New("Missing Auth Token")
  }
  if (phoneNumberID == "") {
    return errors.New("Missing Phone Number ID")
  }
  if (ms.To == "") {
    return errors.New("Missing To Field")
  }
  if (ms.Text.Body == "") {
    return errors.New("Missing Body Field")
  }
  return nil
}

func (ms *MessageSender) send(jsonPayload []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
    return err
	}
	req.Header.Set("Authorization", "Bearer " + authToken)
	req.Header.Set("Content-Type", "application/json") 
	client := &http.Client{}
  resp, err := client.Do(req)
	if err != nil {
    return err
	}
	defer resp.Body.Close()

  return nil
}

func (ms *MessageSender) Send() error {
  payload, err := ms.getPayload()
  if err != nil {
    return err
  }
  err = ms.validate()
  if err != nil {
    return err
  }
  return ms.send(payload)
}
