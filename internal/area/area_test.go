package area

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParseAsAreaLine(t *testing.T) {

  testData := [] struct {
    Input string
    Expected string
  }{
    {
      Input:"I am extre text Filhos de Eva", 
      Expected:"filhos de eva",
    },
    {
      Input: "Mãe Velha stuff at the end", 
      Expected: "mãe velha", 
    },
    {
      Input: "Aruãs", 
      Expected: "aruãs", 
    },
    {
      Input: "Pai AGOSTINHO", 
      Expected: "pai agostinho", 
    },
    {
      Input: "Pai Armando", 
      Expected: "pai armando", 
    },
    {
      Input: "Novo Brasil", 
      Expected: "novo brasil", 
    },
    {
      Input: "China", 
      Expected: "china", 
    },
    {
      Input: "Sta. Lourdes", 
      Expected: "sta. lourdes", 
    },
    {
      Input: " Espirito Santo ", 
      Expected: "espirito santo", 
    },
    {
      Input: "Mucájá",
      Expected: "mucájá",
    },
  }

  for _, test := range testData {
    result, found := ParseAsAreaLine(test.Input)
    assert.True(t, found, "area not found")
    assert.Equal(t, test.Expected, result, "area does not match")
  }
}
