package struct_helper

import (
	"encoding/json"
)

func ToMap(a interface{}) map[string]interface{} {
	var b map[string]interface{}
	inrec, _ := json.Marshal(a)
	json.Unmarshal(inrec, &b)
	return b
}
