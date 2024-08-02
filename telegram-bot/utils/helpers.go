package utils

import (
	"reflect"
	"strings"
)

func ExecuteMethods(obj interface{}) {
	// Calling all the struct methods
	structValue := reflect.ValueOf(obj)
	for i := 0; i < structValue.NumMethod(); i++ {
		method := structValue.Method(i)
		method.Call(nil)
	}
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
