package date

import (
  "fmt"
  "strings"
  "time"
)

// Dates will be detected as:
// dd/mm 
func ParseAsDateLine(line string) (string, bool) {

  var day, month int

  parts := strings.Split(line, " ")
  for _, part := range(parts)  {
    n, err := fmt.Sscanf(part, "%d/%d", &day, &month)

    if err != nil {
      continue
    }

    if n != 2 {
      continue
    }

    // @todo: Can do better than just this, but good enough for now.
    if day > 31 {
      continue
    }

    // @todo: Can do better than just this, but good enough for now.
    if month > 12 {
      continue
    }

    return MonthDayToUTC(month, day), true
  }

  return "", false
}

func MonthDayToUTC(month, day int) string {
  currentYear := time.Now().Year()
  tm := time.Date(currentYear, time.Month(month), day, 0, 0, 0, 0, time.UTC) 
  return tm.Format(time.RFC3339)
}
