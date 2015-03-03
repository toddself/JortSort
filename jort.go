package jort

import "sort"

type Sortable []interface{}

func (s Sortable) Sort() {
	sort.Sort(s)
}

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Sortable) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func JortSort(originalArray Sortable) bool {
	var newArray Sortable
	copy(originalArray, newArray)
	sort.Sort(newArray)

	for i, _ := range newArray {
		if newArray[i] != originalArray[i] {
			return false
		}
	}

	return true
}
