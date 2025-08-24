package textmsg

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestTextMsg(t *testing.T) {
  expected := `{"messaging_product":"whatsapp","type":"text","to":"01234567890","text":{"body":"Some Message"}}`
  ms := NewMessageSender("01234567890", "Some Message")
  bytes, err := ms.getPayload()
  assert.Nil(t, err, "could not get payload")
  assert.Equal(t, expected, string(bytes), "payload does not match")
}
