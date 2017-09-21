package circular

import (
	"reflect"
	"unsafe"
)

type checker struct {
	x unsafe.Pointer
	t reflect.Type
}

func IsCircular(x interface{}) bool {
	return isCircular(reflect.ValueOf(&x).Elem(), make(map[checker]bool))
}

func isCircular(x reflect.Value, seen map[checker]bool) bool {
	if x.CanAddr() {
		ptr := unsafe.Pointer(x.UnsafeAddr())

		c := checker{ptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCircular(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCircular(x.Index(i), seen) {
				return true
			}
		}

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCircular(x.Field(i), seen) {
				return true
			}
		}

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCircular(x.MapIndex(k), seen) {
				return true
			}
		}
	}
	return false
}
