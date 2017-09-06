package labreflect

import (
	"reflect"
)

// Check if 2 objects are equals
func DeepEqual(o1, o2 interface{}) bool {
	return reflect.DeepEqual(o1, o2)
}
