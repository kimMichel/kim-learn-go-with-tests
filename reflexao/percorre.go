package reflexao

import "reflect"

func percorre(x interface{}, fn func(input string)) {
	value := getValue(x)

	percorreValue := func(value reflect.Value) {
		percorre(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			percorreValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			percorreValue(value.Index(i))
		}
	case reflect.String:
		fn(value.String())
	case reflect.Map:
		for _, key := range value.MapKeys() {
			percorreValue(value.MapIndex(key))
		}
	}

}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
