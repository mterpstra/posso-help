package weather

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWeatherGeolocation(t *testing.T) {

  testData := [] struct {
    Input string
    Expected *WeatherResponse
  }{
    {
      Input:"Jensen Beach  FL", 
      Expected:&WeatherResponse{
      },
    },
  }

  for _, test := range testData {
    _, err := GetGeolocation(test.Input)
    assert.Nil(t, err, "error should be nil")
  }
}
