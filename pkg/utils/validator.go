package utils

import (
	"errors"
	"regexp"
	"strings"
)

func PhoneNumberValidator(phone string) bool {
	val := regexp.MustCompile(`^(\+62|62)?[\s-]?0?8[1-9]{1}\d{1}[\s-]?\d{4}[\s-]?\d{2,5}$`)
	return val.MatchString(phone)
}

func IsPhoneNo(s string) bool {
	val := regexp.MustCompile(`^\+?(\d[\d-. ]+)?(\([\d-. ]+\))?[\d-. ]+\d$`)
	return val.MatchString(s)
}

func IsNumeric(s string) bool {
	val := regexp.MustCompile(`^[0-9\s,]*$`)
	return val.MatchString(s)
}

func IsNilOrEmpty(data *string) bool {
	return data == nil || strings.TrimSpace(*data) == ""
}

func DetectDuplicateInt64InSlice(intSlice []int64) error {
	allKeys := make(map[int64]bool)
	list := []int64{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	if len(intSlice) != len(list) {
		return errors.New("Duplicate value detected")
	}
	return nil
}

func RemoveDuplicateInt64InSlice(intSlice []int64) []int64 {
	allKeys := make(map[int64]bool)
	list := []int64{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func CountDuplicateInt64InSlice(list []int64) map[int64]int {

	duplicate_frequency := make(map[int64]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}
