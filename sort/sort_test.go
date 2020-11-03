package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuble(t *testing.T) {
	slice := []int{5, 2, 9, 4, 3, 6, 1}
	if !deepE(buble(slice)) {
		t.Error("failed buble")
		fmt.Println(slice)
	}
}

func TestSelect(t *testing.T) {
	slice := []int{5, 2, 9, 4, 3, 6, 1}
	if !deepE(selectSort(slice)) {
		t.Error("failed select")
		fmt.Println(slice)
	}
}

func TestInsertion(t *testing.T) {
	slice := []int{5, 2, 9, 4, 3, 6, 1}
	if !deepE(insertion(slice)) {
		t.Error("failed select")
		fmt.Println(slice)
	}
}

func TestQuick(t *testing.T) {
	slice := []int{5, 2, 9, 4, 3, 6, 1}
	if !deepE(quick(slice)) {
		t.Error("failed quick")
		fmt.Println(slice)
	}
}

func TestMerge(t *testing.T) {
	slice := []int{5, 2, 9, 4, 3, 6, 1}
	mergeSort(slice)
	fmt.Println(slice)
	if !deepE(slice) {
		t.Error("failed merge")
	}
}

func deepE(result []int) bool {
	expected := []int{1, 2, 3, 4, 5, 6, 9}
	return reflect.DeepEqual(result, expected)
}
