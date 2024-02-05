package utils

import "reflect"

func IsZero(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
