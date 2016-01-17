package defaults

// The method used to access an object's properties.
// The default implementation handles dot notation nesting (i.e. a.b.c).
func DefaultGet(object interface{}, path string) interface{} {
  return nil
}