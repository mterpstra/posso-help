package puretag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewStringSet(
    tag.NewString("fft", []string{"fft", "pure"}),
  )
}
