package main

import (
	"bufio"
	"fmt"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

// 対象のペアを出力する
func symmetric(input [][]int) {
	cache := make(map[int]int)

	for _, v := range input {
		i := v[0]
		j := v[1]
		if cache[j] != i {
			cache[i] = j
			continue
		}
		fmt.Fprintf(out, "%d ", v)
	}
	out.Flush()
}

func pair(i int, j int) []int {
	s := make([]int, 2, 2)
	s[0], s[1] = i, j
	return s
}

func main() {
	input := make([][]int, 0)
	input = append(input, pair(1, 2))
	input = append(input, pair(3, 5))
	input = append(input, pair(4, 7))
	input = append(input, pair(5, 3))
	input = append(input, pair(7, 4))
	symmetric(input)
}
