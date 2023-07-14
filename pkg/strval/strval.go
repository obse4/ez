package strval

import (
	"encoding/json"
	"strconv"
)

// interface 转 string
func Any2String(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value := value.(type) {
	case float64:
		key = strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int:
		key = strconv.Itoa(value)
	case uint:
		key = strconv.Itoa(int(value))
	case int8:
		key = strconv.Itoa(int(value))
	case uint8:
		key = strconv.Itoa(int(value))
	case int16:
		key = strconv.Itoa(int(value))
	case uint16:
		key = strconv.Itoa(int(value))
	case int32:
		key = strconv.Itoa(int(value))
	case uint32:
		key = strconv.Itoa(int(value))
	case int64:
		key = strconv.FormatInt(value, 10)
	case uint64:
		key = strconv.FormatUint(value, 10)
	case string:
		key = value
	case []byte:
		key = string(value)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
