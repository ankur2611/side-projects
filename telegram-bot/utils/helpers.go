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

func ConstructFullName(firstName, lastName, userName string) string {

	name := ""
	if !IsEmpty(firstName) {
		name = firstName
	}
	if !IsEmpty(lastName) {
		name += " " + lastName
	}
	if !IsEmpty(userName) {
		name += " (@" + userName + ")"
	}
	return name
}
