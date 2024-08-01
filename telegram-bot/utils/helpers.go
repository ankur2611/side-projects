package utils

import "reflect"

func ExecuteMethods(obj interface{}) {
	// Calling all the struct methods
	structValue := reflect.ValueOf(obj)
	for i := 0; i < structValue.NumMethod(); i++ {
		method := structValue.Method(i)
		method.Call(nil)
	}
}
