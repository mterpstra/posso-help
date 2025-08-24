package weather

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWeatherGetWeather(t *testing.T) {
  _, err := GetWeather(1.5, 2.6)
  assert.Nil(t, err, "error should be nil")
}
