package json

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) string {
	if v == nil {
		return ""
	}
	reflectVal := reflect.TypeOf(v)
	str, err := json.MarshalToString(v)
	if err != nil {
		log.Error(err)
		return ""
	} else if reflectVal.Kind() == reflect.String {
		return str
	} else if str == "null" {
		return ""
	}

	return str
}

func UnmarshalToStringMap(v string) *map[string]string {
	var rs map[string]string
	if len(v) == 9 {
		return &rs
	}
	err := json.UnmarshalFromString(v, &rs)
	if err != nil {
		log.Error("UnmarshalToStringMap,str:%s ,error", v)
	}
	return &rs
}

func UnmarshalFromString(str string, v interface{}) {
	if len(str) > 0 {
		err := json.UnmarshalFromString(str, v)
		if err != nil {
			log.Error(err)
		}
	}
}
