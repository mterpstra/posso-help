package tag

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestIsInText(t *testing.T) {
  tag := Tag{
    Value: "male",
    Variants: []string{"male", "m"},
  }
  assert.True(t,  tag.IsInText("1111 m angus"))
  assert.True(t,  tag.IsInText("m 1111 angus"))
  assert.True(t,  tag.IsInText("1111 angus m"))
  assert.True(t,  tag.IsInText("1111 male angus"))
  assert.True(t,  tag.IsInText("male 1111 angus"))
  assert.True(t,  tag.IsInText("1111 angus male"))
  assert.True(t,  tag.IsInText("1111   m   angus"))
  assert.True(t,  tag.IsInText("  m 1111 angus"))
  assert.True(t,  tag.IsInText("1111 angus m  "))
  assert.False(t, tag.IsInText("2222 female angus"), "female should return false")
  assert.False(t, tag.IsInText("2222 f angus"))
}

func TestTagList(t *testing.T) {
  tagList := TagList {
    Tags: []Tag { 
      Tag{
        Value: "nelore",
        Variants: []string{"nelore", "nalore"},
      },
      Tag{
        Value: "angus",
        Variants: []string{"angus"},
      },
    },
  }

  found, tag := tagList.IsInText("1111 nelore male")
  assert.True(t, found)
  assert.NotNil(t, tag)
  assert.Equal(t, "nelore", tag.Value)

  found, tag = tagList.IsInText("1111 nalore male")
  assert.True(t, found)
  assert.NotNil(t, tag)
  assert.Equal(t, "nelore", tag.Value)

  found, tag = tagList.IsInText("angus 1111 male")
  assert.True(t, found)
  assert.NotNil(t, tag)
  assert.Equal(t, "angus", tag.Value)

  found, tag = tagList.IsInText("1111 male angus ")
  assert.True(t, found)
  assert.NotNil(t, tag)
  assert.Equal(t, "angus", tag.Value)
}
