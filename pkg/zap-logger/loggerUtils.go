package zaplogger

import (
	"encoding/json"
)

func TransformToString(v interface{}) string {
	var result string

	bytes, err := json.Marshal(v)
	if err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}

	return result
}
