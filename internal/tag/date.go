package tag

import (
  "fmt"
  "log"
  "regexp"
  "strings"
	"time"
)

type Date struct {
}

func NewDate() *Date {
  return &Date{}
}

// Parse accepts a string as input and returns a date if one
// is found.  The date is returned as yyyy-mm-dd
// Example:  dd/mm => yyyy-mm-dd
func (d *Date) Parse(text string) string {
  dateFound := d.findDate(text)
  if dateFound != "" {
    return dateFound
  }
  return d.findDateShort(text)
}

func (d *Date) findDateShort(text string) string {
	dateRegex := regexp.MustCompile(`(\d{1,2})([/-])(\d{1,2})`)
	matches := dateRegex.FindStringSubmatch(text)
  if len(matches) != 4 {
    return ""
  }
  year := time.Now().Year()
  possibleDate := fmt.Sprintf("%04d-%s-%s", year, matches[3], matches[1])
  return d.getFormattedDateIfValid(possibleDate)
}

func (d *Date) findDate(text string) string {
	dateRegex := regexp.MustCompile(`(\d{4})([/-])(\d{1,2})([/-])(\d{1,2})`)
	matches := dateRegex.FindStringSubmatch(text)
  if len(matches) == 0 {
    return ""
  }
  return d.getFormattedDateIfValid(matches[0])
}


func (d *Date) getFormattedDateIfValid(input string) string {
	layouts := []string{
    "2006-01-02",  // YYYY-MM-DD
    "2006-1-2",    // YYYY-M-D
    "2006-01-2",   // YYYY-MM-D
    "2006-1-02",   // YYYY-M-DD
    "01-02-2006",  // MM-DD-YYYY
    "1-02-2006",   // M-DD-YYYY
    "01-2-2006",   // MM-D-YYYY
    "1-2-2006",    // M-D-YYYY
  }
  for _, layout := range layouts {
    layouts = append(layouts, strings.ReplaceAll(layout, "-", "/"))
  }
  for _, layout := range layouts {
    parsedTime, err := time.Parse(layout, input)
    if err == nil {
      return parsedTime.Format("2006-01-02")
    }
  }
  log.Printf("Found possible date: %s, but it was not valid", input)
  return ""
}
