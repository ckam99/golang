package utils

import (
	"encoding/json"
)

func ConvertStructToMap(v *interface{}) map[string]interface{} {
	m, _ := json.Marshal(v)
	var x map[string]interface{}
	_ = json.Unmarshal(m, &x)
	return x
}
