package equalNum

import (
	"fmt"
	"math"
	"reflect"
)

func diff(x, y reflect.Value) (uint64, error) {
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		if (x.Int() > y.Int()) {
			return uint64(x.Int()-y.Int()), nil
		} else if (x.Int() < y.Int()) {
			return uint64(y.Int()-x.Int()), nil
		} else {
			return 0, nil
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		if (x.Uint() > y.Uint()) {
			return x.Uint()-y.Uint(), nil
		} else if (x.Uint() < y.Uint()) {
			return y.Uint()-x.Uint(), nil
		} else {
			return 0, nil
		}

	case reflect.Float32, reflect.Float64:
		if (x.Float() > y.Float()) {
			return math.Float64bits(x.Float() - y.Float()), nil
		} else if (x.Float() < y.Float()) {
			return math.Float64bits(y.Float() - x.Float()), nil
		} else {
			return 0, nil
		}
	
	default:
		return 0, fmt.Errorf("NaN")
	}
}

func equal(x, y reflect.Value) (bool, error) {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid(), nil
	}
	if x.Type() != y.Type() {
		return false, fmt.Errorf("different types")
	}

	d, err := diff(x, y)
	if err != nil {
		return false, fmt.Errorf("13.1: %v", err)
	}

	if float64(d) < 1.0/1000000000.0 {
		return true, nil
	} else {
		return false, nil
	}
}

func Equal(x, y any) (bool, error)  {
	return equal(reflect.ValueOf(x), reflect.ValueOf(y))
}
