package sextag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewStringSet(
    tag.NewString("male", []string{"male", "m"}),
    tag.NewString("female", []string{"female", "f"}),
  )
}
