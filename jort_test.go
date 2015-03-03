package jort

import "testing"

func TestStringSort(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if JortSort(arr) == false {
		t.Error("Your JortSort is broken my friend")
	}

	arr2 := []string{"b", "d", "q"}
	if JortSort(arr2) {
		t.Error("Your JortSort is broken my friend")
	}

}
