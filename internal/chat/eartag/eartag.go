package eartag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewNumber(3,8)
}
