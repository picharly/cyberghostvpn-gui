package tools

import "reflect"

// Return true if the variable is a pointer, and false in the other case
func IsPointer(obj interface{}) bool {
	if obj != nil && reflect.ValueOf(obj).Kind() == reflect.Ptr {
		return true
	}
	return false
}

// Return true if the variable is a slice (even if it is a pointer), and false in the other case
func IsSlice(obj interface{}) bool {
	if obj != nil {
		var k reflect.Kind
		if IsPointer(obj) {
			k = reflect.TypeOf(obj).Elem().Kind()
		} else {
			k = reflect.TypeOf(obj).Kind()
		}
		return (k == reflect.Slice || k == reflect.Array)
	}
	return false
}
