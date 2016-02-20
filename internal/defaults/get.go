package defaults

import "strings"
import "reflect"
import "github.com/antoan-angelov/fuzzy/internal/models"

// The method used to access an object's properties.
// The default implementation handles dot notation nesting (i.e. a.b.c).
func DefaultGet(object interface{}, path string) (interface{}, error) {
	pathComponents := strings.Split(path, ".")

	obj := reflect.ValueOf(object)

	for _, fieldName := range pathComponents {
		immutable := reflect.Indirect(obj)
		obj = immutable.FieldByName(fieldName)
	}

	switch obj.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return obj.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return obj.Uint(), nil
	case reflect.Bool:
		return obj.Bool(), nil
	case reflect.String:
		return obj.String(), nil
	case reflect.Float32, reflect.Float64:
		return obj.Float(), nil
	case reflect.Interface:
		return obj.Interface(), nil
	case reflect.Complex64, reflect.Complex128:
		return obj.Complex(), nil
	}

	return nil, &models.InvalidKeyError{}
}
