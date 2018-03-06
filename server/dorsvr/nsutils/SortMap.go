package utils

import "sort"

type NameAndValue struct {
	key   string
	value interface{}
}

func SortMap(list map[string]interface{}) []*NameAndValue {

	var keyList []string
	var result []*NameAndValue

	for key, _ := range list {
		keyList = append(keyList, key)
	}
	sort.Strings(keyList)
	for _, value := range keyList {
		result = append(result, &NameAndValue{value, list[value]})
	}
	return result
}

func (l *NameAndValue) GetKey() string {
	return l.key
}
func (l *NameAndValue) GetValue() interface{} {

	return l.value
}
