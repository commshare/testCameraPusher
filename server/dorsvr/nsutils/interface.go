package utils

import (
	//"reflect"
	"strconv"
)

func GetType(x interface{}) map[string]interface{} {
	/*reflect.TypeOf(x).String()
	result := make(map[string]interface{})
	switch x.(type){
	case map[string]interface{}:
		{
			j1 := x.(map[string]interface{})
			for key,value := range j1 {

			}
		}
		break
	}*/
	return nil
}

func String2Int(name string) (int64, error) {
	return strconv.ParseInt(name, 10, 64)
}

func Int2String(name int64) string {
	return strconv.FormatInt(name, 10)
}
