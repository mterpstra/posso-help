package eartag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewNumber(4,8)
}
