package jort

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestStringSort(t *testing.T) {
	arr := []string{"a", "b", "c"}
	arrSort := []interface{}{}

	for _, v := range arr {
		arrSort = append(arrSort, v)
	}

	expect(t, JortSort(arrSort), true)

	arr2 := []string{"q", "c", "m"}
	arr2Sort := []interface{}{}

	for _, v := range arr2 {
		arr2Sort = append(arr2Sort, v)
	}

	expect(t, JortSort(arr2Sort), false)
}

func TestIntSort(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	sortableInts := Ints2Sortable(ints)
	expect(t, JortSort(sortableInts), true)

	ints2 := []int{4, 3, 1, 8}
	sortableInts2 := Ints2Sortable(ints2)
	expect(t, JortSort(sortableInts2), false)
}

func TestStringToSortable(t *testing.T) {
	arr := []string{"a", "b", "c"}
	sortableArray := Strings2Sortable(arr)
	for i, _ := range sortableArray {
		expect(t, arr[i], sortableArray[i])
	}
}

func TestIntToSortable(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	sortableArray := Ints2Sortable(arr)
	for i, _ := range sortableArray {
		expect(t, arr[i], sortableArray[i])
	}
}

func TestFloat64ToSortable(t *testing.T) {
	arr := []float64{1.0, 2.0, 3.0}
	sortableArray := Float642Sortable(arr)
	for i, _ := range sortableArray {
		expect(t, arr[i], sortableArray[i])
	}
}
