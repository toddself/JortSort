// A go implementation of http://jort.technology.
// MIT License
package jort

import "sort"

// A generic sortable array that allows us to sort anything we want with only
// one sort method.
type Sortable []interface{}

func (s Sortable) Sort()    { sort.Sort(s) }
func (s Sortable) Len() int { return len(s) }
func (s Sortable) Less(i, j int) bool {
	if _, ok := s[i].(float64); ok {
		return s[i].(float64) < s[j].(float64)
	}

	if _, ok := s[i].(int); ok {
		return s[i].(int) < s[j].(int)
	}

	if _, ok := s[i].(string); ok {
		return s[i].(string) < s[j].(string)
	}

	return false
}

func (s Sortable) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Checks to see if an array is already sorted. Returns `true` if it's sorted
// `false` if it's not.
func JortSort(originalArray Sortable) bool {
	newArray := Sortable{}
	for _, v := range originalArray {
		newArray = append(newArray, v)
	}
	sort.Sort(newArray)

	for i, _ := range newArray {
		if newArray[i] != originalArray[i] {
			return false
		}
	}

	return true
}

// Converts `[]string` to `Sortable` so it can be passed into `JortSort`
func Strings2Sortable(str []string) Sortable {
	arr := []interface{}{}
	for _, v := range str {
		arr = append(arr, v)
	}
	return arr
}

// Converts `[]int` to `Sortable`
func Ints2Sortable(ints []int) Sortable {
	arr := []interface{}{}
	for _, v := range ints {
		arr = append(arr, v)
	}
	return arr
}

// Converts `[]float64` to `Sortable`
func Float642Sortable(ints []float64) Sortable {
	arr := []interface{}{}
	for _, v := range ints {
		arr = append(arr, v)
	}
	return arr
}
