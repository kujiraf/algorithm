package main

import (
	"fmt"
	"strconv"
)

func main() {
	typeCheck()
	recursiveAnonymousFunc()
	hash("abc", 7)
}

func hash(str string, size int) {
	sum := 0
	for _, v := range str {
		sum += int(v)
	}

	fmt.Printf("sum=%d, hash=%d\n", sum, sum%size)
}

func typeCheck() {
	var i interface{}
	var j interface{}
	i = 1
	j = "1"

	k, b := i.(int)
	fmt.Printf("isInt:%v, %d\n", b, k)

	l, b := j.(int)
	fmt.Printf("isInt:%v, %d\n", b, l)
}

func recursiveAnonymousFunc() {
	count := 0

	var anoF func(x []string) []string
	anoF = func(x []string) []string {
		if count >= 10 {
			return x
		}
		count++

		str := "x" + strconv.Itoa(count)
		return anoF(append(x, str))
	}

	fmt.Println(anoF([]string{"x" + strconv.Itoa(count)}))
}
