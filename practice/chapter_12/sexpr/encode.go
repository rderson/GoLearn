// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// !+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}, json bool) ([]byte, error) {
	var buf bytes.Buffer
	if json {
		if err := encodeJSON(&buf, reflect.ValueOf(v)); err != nil {
			return nil, err
		}
	} else {
		if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
// !+encode
func encode(buf *bytes.Buffer, v reflect.Value, depth int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), depth)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		first := true
		for i := 0; i < v.Len(); i++ {
			if checkNil(v.Index(i)) {
				continue
			}
			if !first {
				buf.WriteByte(' ')
			}
			first = false
			if err := encode(buf, v.Index(i), depth); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte('\n')
				if i == 0 {
					depth++
				}
				for i := 0; i < depth; i++ {
					buf.WriteByte('\t')
				}
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		first := true
		for i := 0; i < v.NumField(); i++ {
			if checkNil(v.Field(i)) {
				continue
			}
			if !first {
				buf.WriteByte(' ')
			}
			first = false
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), depth); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i != v.NumField()-1 {
				buf.WriteByte('\n')
				for i := 0; i < depth; i++ {
					buf.WriteByte('\t')
				}
			}
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		first := true
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if checkNil(key) && checkNil(v.MapIndex(key)) {
				continue
			}
			if !first {
				buf.WriteByte(' ')
			}
			first = false
			buf.WriteByte('(')
			if err := encode(buf, key, depth); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), depth); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i != v.Len()-1 {
				buf.WriteByte('\n')
				if i == 0 {
					depth++
				}
				for i := 0; i < depth; i++ {
					buf.WriteByte('\t')
				}
			}
		}
		buf.WriteByte(')')

	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%0.1f %0.1f)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		fmt.Fprintf(buf, "%q %v", v.Type().String(), v.Elem())

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func encodeJSON(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encodeJSON(buf, v.Elem())
	
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteByte(',')
			}
			if err := encodeJSON(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Map:
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i != 0 {
				buf.WriteByte(',')
			}
			if err := encodeJSON(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encodeJSON(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			if err := encodeJSON(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Interface:
		if v.IsNil() {
			buf.WriteString("null")
			return nil
		}
		return encodeJSON(buf, v.Elem())

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	
	}

	return nil
}

func checkNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return true
		}
		return false

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() == 0 {
			return true
		}
		return false

	case reflect.String:
		if v.String() == "" {
			return true
		}
		return false

	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return false

	case reflect.Slice: // (value ...)
		if !v.IsNil() {
			return false
		}
		return true
		
	case reflect.Struct: // ((name value) ...)
		for i := 0; i < v.NumField(); i++ {
			if !checkNil(v.Field(i)) {
				return false
			}
		}
		return true

	case reflect.Map: // ((key value) ...)
		if v.IsNil() {
			return true
		}
		return false

	case reflect.Bool:
		if v.Bool() {
			return true
		}
		return false

	case reflect.Float32, reflect.Float64:
		if v.Float() == 0.0 {
			return true
		}
		return false
	
	case reflect.Complex64, reflect.Complex128:
		if v.Complex() == 0 {
			return true
		}
		return false

	case reflect.Interface:
		if v.IsNil() {
			return true
		}
		return false

	default: // chan, func, interface
		return true
	}
}
//!-encode
