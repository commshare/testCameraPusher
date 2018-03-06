package utils

import "strconv"

func Convert2String(value interface{}) string {
	switch value.(type) {
	case int:
		{
			return strconv.FormatInt(int64(value.(int)), 10)
		}
		break
	case int64:
		{
			return strconv.FormatInt(value.(int64), 10)
		}
		break
	case float32:
		{
			return strconv.FormatFloat(float64(value.(float32)), 'f', 0, 32)
		}
		break
	case float64:
		{
			return strconv.FormatFloat(value.(float64), 'f', 0, 64)
		}
		break
	case string:
		{
			return value.(string)
		}
		break
	case []byte:
		{
			return string(value.([]byte))
		}
		break
	}
	return ""
}
