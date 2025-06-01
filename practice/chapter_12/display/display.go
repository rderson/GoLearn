// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 333.

// Package display provides a means to display structured data.
package display

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const maxDepth = 10

//!+Display

func Display(name string, x interface{}) {
	currentRecursion := 0
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), currentRecursion, maxDepth)
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func formatComplexKey(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		var parts []string
		for i:=0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Name 
			value := formatComplexKey(v.Field(i))
			parts = append(parts, fmt.Sprintf("%s: %s", field, value))
		}
		return v.Type().Name() + "{" + strings.Join(parts, ", ") + "}"
	case reflect.Array:
		var values []string
		for i := 0; i < v.Len(); i++ {
			values = append(values, formatComplexKey(v.Index(i)))
		}
		return "{" + strings.Join(values, ", ") + "}"
	default:
		return formatAtom(v)
	}
}

// !+display
func display(path string, v reflect.Value, currentDepth int, maxDepth int) {

	if currentDepth > maxDepth {
		fmt.Printf("%s = <max depth reached>\n", path)
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), currentDepth+1, maxDepth)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), currentDepth+1, maxDepth)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatComplexKey(key)), v.MapIndex(key), currentDepth+1, maxDepth)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), currentDepth+1, maxDepth)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), currentDepth+1, maxDepth)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

//!-display
