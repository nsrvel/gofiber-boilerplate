package utils

import "strings"

func GetDBNameFromDriverSource(str string) string {
	if str == "" {
		return ""
	}
	if strings.Contains(str, "dbname=") {
		array1 := strings.Split(str, "dbname=")
		if len(array1) == 0 {
			return ""
		}
		array2 := strings.Split(array1[1], " ")
		if len(array2) == 0 {
			return ""
		}
		return array2[0]
	}
	if strings.Contains(str, "database=") {
		array1 := strings.Split(str, "database=")
		if len(array1) == 0 {
			return ""
		}
		array2 := strings.Split(array1[1], ";")
		if len(array2) == 0 {
			return ""
		}
		return array2[0]
	}
	return ""
}
