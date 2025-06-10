package cycle

import (
	"reflect"
	"unsafe"
)

func IsCyclic(x interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return visit(reflect.ValueOf(x), seen)
}

func visit(v reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !v.IsValid() {
		return false
	}

	switch v.Kind() {

	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return false
		}

		ptr := unsafe.Pointer(v.Pointer())
		if seen[ptr] {
			return true
		} 
		seen[ptr] = true
		return visit(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if visit(v.Index(i), seen) {
				return true
			}
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if visit(v.Field(i), seen) {
                return true
            }
		}
	
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if visit(key, seen) || visit(v.MapIndex(key), seen) {
				return true
			}
		}
	}
	
	return false
}