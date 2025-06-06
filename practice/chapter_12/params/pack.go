// Exercise 12.11: Write the corresponding Pack function. Given a struct value, Pack should
// return a URL incorporating the parameter values from the struct.

package params

import (
	"reflect"
	"strings"
)

func Pack(url string, v interface{}) string {
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
		url += "?"
	}

	var first = true

	data := reflect.ValueOf(v)
	for i := 0; i < data.NumField(); i++ {
		fieldInfo := data.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		
		field := data.Field(i)

		if field.Kind() == reflect.Slice {
			for i := 0; i < field.Len(); i++ {
				fieldValue := convertString(field.Index(i));
				if !first {
					url += "&" + name + "=" + fieldValue
				} else {
					first = false
					url += name + "=" + fieldValue
				}
			}
		} else {
			fieldValue := convertString(field);
			if !first {
				url += "&" + name + "=" + fieldValue
			} else {
				first = false
				url += name + "=" + fieldValue
			}
		}
	}

	return url
}


