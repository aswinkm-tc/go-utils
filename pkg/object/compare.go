package object

import (
	"reflect"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IsEqual compares two kubernetes client.Objects and returns true if they are equal
func IsEqual(old, new client.Object) (bool, error) {
	mergedObject, err := mergeObject(old, new)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(old, mergedObject), nil
}
