package breedtag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewStringSet(
    tag.NewString("angus",  []string{"angus"}),
    tag.NewString("nelore", []string{"nelore", "nalore"}),
  )
}
