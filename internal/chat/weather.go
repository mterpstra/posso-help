package chat

import (
  "fmt"
  "strings"

  "posso-help/internal/utils"
  "posso-help/internal/weather"
)

const REPLY_WEATHER = "Posso Help Weather:\n%s"

type WeatherMessage struct {
  searchAddress string
  displayAddress string
  weather string
}

func (w *WeatherMessage) Parse(message string) bool {
  found := false
  address := ""

  lines := strings.Split(message, "\n")
  for _, line := range lines {
    line = utils.SanitizeLine(line)
    words := strings.Fields(line)
    for _, word := range words {

      if word == "for" || word == "in"  {
        continue
      }

      if found {
        address += word + "+"
      }

      if word == "weather" {
        found = true
      }
    }

    if found && len(address) > 0 {
       break;
    }
  }

  address = strings.TrimRight(address, "+")
  w.searchAddress = address
  return found 
}

func (w *WeatherMessage) Text() string {
  return fmt.Sprintf(REPLY_WEATHER, w.weather)
}

// Acrtually gets the weather for the passed address
func (w *WeatherMessage) Insert(bmv *BaseMessageValues) error {

  geoLocation, err := weather.GetGeolocation(w.searchAddress)
  if err != nil {
    return err
  }

  weather, err := weather.GetWeather(
    geoLocation.Results[0].Geometry.Location.Lat,
    geoLocation.Results[0].Geometry.Location.Lng,
  )
  if err != nil {
    return err
  }

  w.weather = fmt.Sprintf(
    "Address: %s\nTime: %.16s\nCondition: %s\nTemperature: %3.1f %s\nPrecipitation: %2.0f%% %s",
    geoLocation.Results[0].FormattedAddress,
    weather.CurrentTime,
    weather.WeatherCondition.Description.Text,
    weather.Temperature.Degrees,
    utils.Capitalize(weather.Temperature.Unit),
    weather.Precipitation.Probability.Percent,
    utils.Capitalize(weather.Precipitation.Probability.Type))

  return nil
}
