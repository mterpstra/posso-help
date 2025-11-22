package main	


import (
  "fmt"
  "log"
  "sort"
  "strings"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertBsonToCsv(data []bson.D) (string, error) {
  headers := []string{}
  rowValues := map[string]string{}
  results := ""
	for index, doc := range data {
    log.Printf("parsing row: %v\n", doc)
    row := ""
    for _, value := range doc {
      // {name mark}
      nameAndValue := fmt.Sprintf("%v", value)
      nameAndValue = strings.TrimLeft(nameAndValue, "{")
      nameAndValue = strings.TrimRight(nameAndValue, "}")
      parts := strings.SplitN(nameAndValue, " ", 2)
      if parts[0] == "account" || parts[0] == "_id" {
        continue
      }
      if index == 0 {
        headers = append(headers, parts[0])
      }
      rowValues[parts[0]] = parts[1]
    }

    if index == 0 {
      sort.Strings(headers)
    }
    for _, header := range headers {
      row += rowValues[header] + ","
    }
    row = strings.TrimSuffix(row, ",") + "\n"
    log.Printf("Adding row: %v\n", row)
    results += row
  }
  return fmt.Sprintf("%s\n%s", strings.Join(headers, ","), results), nil
}
