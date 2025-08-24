package weather 

import (
  "io"
  "os"
  "fmt"
  "net/http"
  "encoding/json"
)

func GetWeather(latitude, longitude float64) (*WeatherResponse, error) {

  apiKey := os.Getenv("GOOGLE_API_KEY")
  url := os.Getenv("WEATHER_URL")
  fullURL := fmt.Sprintf("%s?location.latitude=%f&location.longitude=%f&key=%s", 
                         url, latitude, longitude, apiKey)

  resp, err := http.Get(fullURL)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close() 

  
  bodyBytes, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  response := &WeatherResponse{}
  err = json.Unmarshal(bodyBytes, &response)
  if err != nil {
    return nil, err
  }

  return response, nil
}
