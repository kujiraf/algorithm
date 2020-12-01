package list

import (
	"fmt"
	"testing"
)

var list = List{}

func setup() {
	fmt.Println("setup list")
	list = List{}
}

func TestAppend(t *testing.T) {
	setup()
	fmt.Println("test appendTest")
	list.Append(&Node{data: 1})
	list.Append(&Node{data: 2})
	list.Append(&Node{data: 3})
	list.Append(&Node{data: "aaa"})
	list.Append(&Node{data: "bbb"})
	list.Append(&Node{data: "ccc"})
	expected := "[1 2 3 aaa bbb ccc]"
	if actual := list.ToString(); actual != expected {
		t.Errorf("got %s, want %s", actual, expected)
	}
}

func TestPrepend(t *testing.T) {
	setup()
	fmt.Println("test appendTest")
	list.Prepend(&Node{data: 1})
	list.Append(&Node{data: 2})
	list.Append(&Node{data: 3})
	list.Prepend(&Node{data: "aaa"})
	list.Append(&Node{data: "bbb"})
	list.Prepend(&Node{data: "ccc"})
	expected := "[ccc aaa 1 2 3 bbb]"
	if actual := list.ToString(); actual != expected {
		t.Errorf("got %s, want %s", actual, expected)
	}
}
