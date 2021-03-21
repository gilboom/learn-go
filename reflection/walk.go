package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	//if val.Kind() == reflect.Slice {
	//	for i := 0; i < val.Len(); i++ {
	//		Walk(val.Index(i).Interface(), fn)
	//	}
	//	return
	//}
	//
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//
	//	//if field.Kind() == reflect.String {
	//	//	fn(field.String())
	//	//}
	//	//
	//	//if field.Kind() == reflect.Struct {
	//	//	Walk(field.Interface(), fn)
	//	//}
	//
	//	switch field.Kind() {
	//	case reflect.String:
	//		fn(field.String())
	//	case reflect.Struct:
	//		Walk(field.Interface(), fn)
	//	}
	//}

	walkValue := func(v reflect.Value) {
		Walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walkValue(val.MapIndex(k))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}

	return val
}
