package util

import "sort"

func CheckInArray(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}
