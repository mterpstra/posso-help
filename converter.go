package main	


import (
  "fmt"
  "strings"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertBsonToCsv(data []bson.D) (string, error) {
  headers := ""
  results := ""
	for index, doc := range data {
    row := ""
    for _, value := range doc {
      // {name mark}
      nameAndValue := fmt.Sprintf("%v", value)
      nameAndValue = strings.TrimLeft(nameAndValue, "{")
      nameAndValue = strings.TrimRight(nameAndValue, "}")
      parts := strings.SplitN(nameAndValue, " ", 2)
      if parts[0] != "_id" {
        row += parts[1] + ","
        if index == 0 {
          headers += parts[0] + ","
        }
      }
    }
    row = strings.TrimSuffix(row, ",")
    row += "\n"
    results += row
  }
  headers = strings.TrimSuffix(headers, ",")
  return fmt.Sprintf("%s\n%s", headers, results), nil
}
