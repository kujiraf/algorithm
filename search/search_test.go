package search

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 9}
	result := linearSearch(slice, 6)
	if result != 5 {
		t.Error("failed linearSearch")
	}
}

func TestBinarySearch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 9}
	result := binarySearch(slice, 6)
	if result != 5 {
		fmt.Println(result)
		t.Error("failed binarySearch")
	}
}

func TestRecBinarySearch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 9}
	result := recursiveBinarySearch(slice, 6, 0, len(slice)-1)
	if result != 5 {
		fmt.Println(result)
		t.Error("failed binarySearch")
	}
}
