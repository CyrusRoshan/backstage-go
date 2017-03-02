package utils

import "encoding/json"

func MustMarshal(data interface{}) []byte {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return out
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
