package deathtag

import (
  "posso-help/internal/chat/tag"
)

func New() tag.Tag {
  return tag.NewStringSet(
    tag.NewString("aborto",       []string{"aborto"}),
    tag.NewString("morreu",       []string{"morreu"}),
    tag.NewString("natimorto",    []string{"natimorto"}),
    tag.NewString("natimortos",   []string{"natimortos"}),

    // Order here matters, morto is a subset of this entry and there is a space
    tag.NewString("nasceu-morto", []string{"nasceu morto"}),
    tag.NewString("morto",        []string{"morto"}),
  )
}

