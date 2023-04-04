package reflexao

import "reflect"

func percorre(x interface{}, fn func(input string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			percorre(field.Interface(), fn)
		}
	}
}
