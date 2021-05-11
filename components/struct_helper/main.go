package struct_helper

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

func ToMap(a interface{}) map[string]interface{} {
	var b map[string]interface{}
	inrec, _ := json.Marshal(a)
	json.Unmarshal(inrec, &b)
	return b
}

func ToStruct(a interface{}) interface{} {
	var result interface{}
	// Decode(a, &result)
	mapstructure.Decode(a, &result)
	return result
}
