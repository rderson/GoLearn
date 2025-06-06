package params

import (
	"reflect"
	"strconv"
)

func convertString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()

	case reflect.Int:
		i := strconv.FormatInt(v.Int(), 10)
		return i

	case reflect.Bool:
		if v.Bool() {
			return "true"
		} 
		return "false"

	default:
		return ""
	}
}

func convertInt(v reflect.Value) int {
	switch v.Kind() {
	case reflect.String:
		i, _ := strconv.Atoi(v.String())
		return i
	
	case reflect.Int:
		return int(v.Int())
	
	case reflect.Bool:
		if v.Bool() {
			return 1
		}
		return 0

	default:
		return 0
	}
}