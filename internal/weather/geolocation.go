package weather 

import (
  "regexp"
  "strings"
  "errors"
  "io"
  "os"
  "fmt"
  "net/http"
  "encoding/json"
)

func inputToAddress(input string) string {
  re := regexp.MustCompile(`\s+`)
  address := re.ReplaceAllString(input, " ")
  address = strings.TrimSpace(address)
  address = strings.ReplaceAll(address, " ", "+")
  // @todo: There must be a better fix for this.
  address = strings.ReplaceAll(address, "ã", "a")
  return address 
}

func GetGeolocation(input string) (*GeolocationResponse, error) {

  apiKey := os.Getenv("GOOGLE_API_KEY")
  url := os.Getenv("GEOLOC_URL")
  address := inputToAddress(input) 
  fullURL := fmt.Sprintf("%s?address=%s&key=%s", url, address, apiKey) 

  resp, err := http.Get(fullURL)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close() 

  bodyBytes, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  response := &GeolocationResponse{}
  // @todo: checks status, if not OK we will crash later
  err = json.Unmarshal(bodyBytes, &response)
  if err != nil {
    return nil, err
  }

  if len(response.Results) < 1 {
    return nil, errors.New("no geolocation results") 
  }

  return response, nil
}

