package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

func GetStructAttrs(obj struct{}) []string {
	attrs := []string{}
	e := reflect.ValueOf(&obj).Elem()
	for i := 0; i < e.NumField(); i++ {
		//varName := e.Type().Field(i).Name
		// varType := e.Type().Field(i).Type
		// varValue := e.Field(i).Interface()
		// fmt.Printf("%v %v %v\n", varName, varType, varValue)
		attrs = append(attrs, strings.ToLower(e.Type().Field(i).Name))
	}
	return attrs
}

func GetStructAttrsValues(obj interface{}) map[string]interface{} {
	attrs := make(map[string]interface{})
	e := reflect.ValueOf(&obj).Elem()
	for i := 0; i < e.NumField(); i++ {
		attrs[strings.ToLower(e.Type().Field(i).Name)] = e.Field(i).Interface()
	}
	return attrs
}

func StructToMap(v *interface{}) map[string]interface{} {
	m, _ := json.Marshal(v)
	var x map[string]interface{}
	_ = json.Unmarshal(m, &x)
	return x
}
