package main

import (
	"fmt"
	"strconv"
)

func main() {
	// typeCheck()
	// recursiveAnonymousFunc()
	// hash("abc", 7)

	v := v{data: 5}
	fmt.Printf("main0: %p, %d\n", &v, v.data)
	v.print1()
	fmt.Printf("main1: %p, %d\n", &v, v.data)
	v.print2()
	fmt.Printf("main2: %p, %d\n", &v, v.data)
}

type v struct {
	data int
}

func (p *v) print1() {
	fmt.Printf("print1: %p, %T\n", p, p)
	p.data = 4 // ポインタレシーバのため、ここの変更は呼び出し元に影響がある

	var x func(b *v) *v
	x = func(b *v) *v {
		fmt.Printf("print2: %p, %T\n", b, b)
		// b = &v{data: 2} // ここでポインタが変わるので、pのアドレスの値は変更されない
		*b = v{data: 2} // ここはbの値に対して、値を代入しているため、pのアドレスの値（=bの値）も変更される
		fmt.Printf("print3: %p, %T\n", b, b)
		return b
	}

	p = x(p)
	fmt.Printf("print4: %p, %T\n", p, p)
}

func (p v) print2() {
	fmt.Printf("print2: %p, %T\n", &p, p)
	p.data = 3 // 値レシーバのため、ここの変更は呼び出し元には影響しない
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
