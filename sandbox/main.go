package main

import (
	"fmt"
	"strconv"
)

func main() {
	typeCheck()
	recursiveAnonymousFunc()
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

	fmt.Print(anoF([]string{"x" + strconv.Itoa(count)}))
}
