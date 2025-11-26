package tag

type Tag interface {
  Parse(string) bool 
  Value() string
  ValueAsInt() int
}
