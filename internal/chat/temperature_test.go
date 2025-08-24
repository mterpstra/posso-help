package chat

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestTemperatureMessage(t *testing.T) {
  tm := &TemperatureMessage{}
  input := "02/04 80c\n25/12 75C\nanystring\n30/01 40Cans"
  assert.Equal(t, tm.Parse(input), true, "Could not parse temperature message")
  assert.Equal(t, len(tm.Entries), 2, "Wrong number of temperature entries")
  assert.Equal(t, tm.Entries[0].Temperature, 80, "Wrong temperature value")
  assert.Equal(t, tm.Entries[0].Date, "2025-04-02T00:00:00Z", "Wrong temperature date")
  assert.Equal(t, tm.Entries[1].Temperature, 75, "Wrong temperature value")
  assert.Equal(t, tm.Entries[1].Date, "2025-12-25T00:00:00Z", "Wrong temperature date")
}
