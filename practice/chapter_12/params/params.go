// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func isValidEmail(s string) bool {
	matched, _ := regexp.MatchString(`^[^@]+@[^@]+\.[^@]+$`, s)
    return matched
}

func isValidZipCode(i int) bool {
	return i >= 10000 && i <= 99999
}

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
	
		fields[name] = v.Field(i)
	}	

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}

	// Check valid requirements
	for i := 0; i < v.NumField(); i++ {
		fieldStr := convertString(v.Field(i))
		fieldInt := convertInt(v.Field(i))

		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		valid := tag.Get("valid")
		if valid == "email" {
			if !isValidEmail(fieldStr)  {
				return fmt.Errorf("invalid email: %s", fieldStr)
			}
		} else if valid == "zip" {
			if !isValidZipCode(fieldInt) {
				return fmt.Errorf("invalid ZIP code: %d", fieldInt)
			}
		}
	}
	return nil
}

//!-Unpack

// !+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
