package password

import (
  "fmt"
  "errors"
  "os"
  "crypto/md5"
	"encoding/hex"
)

func GetSalted(password string) (string, error) {
  salt := os.Getenv("SALT")
  if len(salt) < 10 {
    return "", errors.New("invalid salt value")
  }
  salted := fmt.Sprintf("%s+%s", salt, password)
	hasher := md5.New() 
	hasher.Write([]byte(salted))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}


