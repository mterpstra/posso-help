package breedtag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewStringSet(
    tag.NewString("angus",     []string{"angus"}),
    tag.NewString("nelore",    []string{"nelore", "nalore"}),
    tag.NewString("brangus",   []string{"brangus"}),
    tag.NewString("sta.zelia", []string{"sta.zelia", "sta. zelia"}),
    tag.NewString("cruzada",   []string{"cruzada"}),
    tag.NewString("cruzado",   []string{"cruzado"}),
  )
}
