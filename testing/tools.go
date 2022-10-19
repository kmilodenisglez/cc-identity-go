package testing

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"strings"
)

// MarshalProtoOrPanic is a helper for proto marshal.
func MarshalProtoOrPanic(pb proto.Message) []byte {
	data, err := proto.Marshal(pb)
	if err != nil {
		panic(err)
	}

	return data
}

func UnmarshalJSONOrPanic(data []byte, v interface{}) interface{} {
	err := json.Unmarshal(data, &v)
	if err != nil {
		panic(err)
	}

	return v
}

func MarshalJSONOrPanic(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

// Dummy implementation to create compose key
func CreateComposeKey(objectType string, keys []string) (string, error) {
	return objectType + "_" + strings.Join(keys, "_"), nil
}

func GomegaString() string {
	return strings.Repeat("randomstring", 10*2)
}
