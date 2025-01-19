package object

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// mergeObject merges two kubernetes client.Objects using the json-patch library
func mergeObject(old, new client.Object) (runtime.Object, error) {
	oldbytes, err := json.Marshal(old)
	if err != nil {
		return nil, err
	}
	newbytes, err := json.Marshal(new)
	if err != nil {
		return nil, err
	}

	patch, err := jsonpatch.CreateMergePatch(oldbytes, newbytes)
	if err != nil {
		return nil, err
	}

	outBytes, err := jsonpatch.MergePatch(oldbytes, patch)
	if err != nil {
		return nil, err
	}

	out := old.DeepCopyObject()
	if err := json.Unmarshal(outBytes, out); err != nil {
		return nil, err
	}

	fmt.Println(out)
	return out, nil
}
