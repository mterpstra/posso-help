package chat

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWeatherMessage(t *testing.T) {

  testdata := []string{
    `Get me the weather for Jupiter FL`,
  }

  bmv := &BaseMessageValues{}
  wm := &WeatherMessage{}

  for _, test := range testdata {
    println("testing", test)
    assert.True(t, wm.Parse(test), "Could not parse weather message")
    assert.Nil(t, wm.Insert(bmv),  "Could not insert weather message")
    assert.Contains(t, wm.Text("en-US"), "Celsius", "Expected string to contain 'celcius'")
    println(wm.Text("en-US"))
  }

}
