package tag

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
  "regexp"
)

func TestIsVariantFound(t *testing.T) {
  tag := Tag{
    Value: "Some-Value",
    Variants: []string{"some-value", "some value", "some val"},
  }
  assert.True(t, tag.IsVariantFound("some-value"))
  assert.True(t, tag.IsVariantFound("some value"))
  assert.True(t, tag.IsVariantFound("some val"))
  assert.False(t, tag.IsVariantFound("false value"))
  assert.False(t, tag.IsVariantFound(" some-value "))
}

func TestToRegexPattern(t *testing.T) {
  tag := Tag{
    Value: "Some-Value",
    Variants: []string{"some-value", "some value", "some val"},
  }
  assert.Equal(t, "(some-value|some value|some val)", 
               tag.ToRegexPattern(), "pattern does not match")
}
