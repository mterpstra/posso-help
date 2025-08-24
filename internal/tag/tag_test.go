package tag 

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestTag(t *testing.T) {
  tag := Tag{}
  err := tag.Find("16166100305", "area")
  assert.Nil(t, err)
  assert.Equal(t, "16166100305", tag.Account)
  assert.Equal(t, "area", tag.Name)
  assert.Equal(t, "area", tag.Name)
  assert.Equal(t, "string", tag.Type)
  assert.Equal(t, 10, len(tag.Values))
}
